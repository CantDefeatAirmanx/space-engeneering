package closer

import "context"

type Closer interface {
	CloseAll(ctx context.Context) error
	Add(funcs ...func(ctx context.Context) error)
	AddNamed(name string, f func(ctx context.Context) error)
}
