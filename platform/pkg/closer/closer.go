package closer

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	"go.uber.org/zap"
)

var _ Closer = (*CloserImpl)(nil)

type CloserImpl struct {
	mu      sync.Mutex
	once    sync.Once
	funcs   []closerFunc
	signals []os.Signal
	logger  Logger
	done    bool
}

type closerFunc struct {
	name string
	fn   func(ctx context.Context) error
}

func NewCloser(ctx context.Context, opts ...CloserOptsFunc) *CloserImpl {
	options := &CloserOpts{
		Logger: logger.NoopLogger(),
		Signals: []os.Signal{
			syscall.SIGINT,
			syscall.SIGTERM,
		},
	}
	for _, opt := range opts {
		opt(options)
	}

	closer := &CloserImpl{
		logger:  options.Logger,
		signals: options.Signals,
	}

	go func() {
		if err := closer.watchSignals(ctx); err != nil {
			closer.logger.Error("Closer compose closing error", zap.Error(err))
		}
	}()

	return closer
}

func (c *CloserImpl) Add(funcs ...func(ctx context.Context) error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, fn := range funcs {
		c.funcs = append(c.funcs, closerFunc{
			name: "",
			fn:   fn,
		})
	}
}

func (c *CloserImpl) AddNamed(name string, f func(ctx context.Context) error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.funcs = append(c.funcs, closerFunc{
		name: name,
		fn:   f,
	})
}

func (c *CloserImpl) CloseAll(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.done {
		c.logger.Error("Second call to Closer.CloseAll, ignoring")
		return nil
	}

	result := make([]error, 0, len(c.funcs))

	c.once.Do(func() {
		wg := sync.WaitGroup{}
		errCh := make(chan error)

		for _, fn := range c.funcs {
			wg.Add(1)

			go func() {
				defer wg.Done()
				defer func() {
					if r := recover(); r != nil {
						c.logger.Error(fmt.Sprintf("Panic while closing %s - %v", c.getFuncTitle(fn), r))
						if err, ok := r.(error); ok {
							errCh <- err
						}
					}
				}()

				err := c.close(ctx, fn)
				if err != nil {
					errCh <- err
				}
			}()
		}

		go func() {
			wg.Wait()
			close(errCh)
		}()

		for {
			select {
			case <-ctx.Done():
				c.logger.Error("Context done before all closer functions closed")
				result = append(result, ctx.Err())
				return
			case err, isOpen := <-errCh:
				if !isOpen {
					c.logger.Info("All closer functions successfully closed")
					return
				}
				result = append(result, err)
			}
		}
	})
	c.done = true

	return errors.Join(result...)
}

func (c *CloserImpl) close(ctx context.Context, fn closerFunc) error {
	err := fn.fn(ctx)
	if err != nil {
		c.logger.Error(fmt.Sprintf("Error while closing %s - %v", c.getFuncTitle(fn), err))
	}
	return err
}

func (c *CloserImpl) getFuncTitle(fn closerFunc) string {
	if fn.name == "" {
		return "Unnamed closer function"
	}
	return fmt.Sprintf("Closer function %s", fn.name)
}

func (c *CloserImpl) watchSignals(ctx context.Context) error {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, c.signals...)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case signal := <-signals:
			c.logger.Info(fmt.Sprintf("Received signal: %s", signal.String()))
			c.CloseAll(ctx)
		}
	}
}
