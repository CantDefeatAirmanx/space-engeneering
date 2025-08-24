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
	funcs   []closerFunc
	signals []os.Signal
	logger  Logger
	done    chan status
	status  status
}

type status int

const (
	StatusIdle status = iota
	StatusRunning
	StatusDone
)

type closerFunc struct {
	name string
	fn   func(ctx context.Context) error
}

func NewCloser(ctx context.Context, opts ...CloserOptsFunc) (*CloserImpl, chan status) {
	options := &CloserOpts{
		Logger: logger.DefaultInfoLogger(),
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
		done:    make(chan status),
		status:  StatusIdle,
	}

	go func() {
		if err := closer.watchSignals(ctx); err != nil {
			closer.logger.Error("Closer compose closing error", zap.Error(err))
		}
	}()

	return closer, closer.done
}

func (c *CloserImpl) SetLogger(logger Logger) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.logger = logger
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
	if c.status == StatusRunning {
		return nil
	}
	if c.status == StatusDone {
		c.logger.Info("Closer already done")
		return nil
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.status = StatusRunning

	result := make([]error, 0, len(c.funcs))

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

outer:
	for {
		select {
		case <-ctx.Done():
			c.logger.Error("Context done before all closer functions closed")
			result = append(result, ctx.Err())
			break outer
		case err, isOpen := <-errCh:
			if !isOpen {
				c.logger.Info("All closer functions successfully closed")
				break outer
			}
			result = append(result, err)
		}
	}
	c.status = StatusDone
	c.done <- StatusDone
	close(c.done)

	return errors.Join(result...)
}

func (c *CloserImpl) close(ctx context.Context, fn closerFunc) error {
	c.logger.Info(fmt.Sprintf("Closing %s", c.getFuncTitle(fn)))
	err := fn.fn(ctx)

	if err != nil {
		c.logger.Error(fmt.Sprintf("Error while closing %s - %v", c.getFuncTitle(fn), err))
		return err
	}
	c.logger.Info(fmt.Sprintf("Successfully closed %s", c.getFuncTitle(fn)))

	return nil
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
			return c.CloseAll(ctx)
		}
	}
}
