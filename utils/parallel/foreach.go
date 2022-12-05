package parallel

import (
	"context"

	"golang.org/x/sync/errgroup"
)

// ForEachString is a utility wrapper to help with pointer-izing a string (common scenario)
// WARNING: handler is not threadsafe. You may need to lock non-threadsafe resources or use other techniques
func ForEachString(ctx context.Context, workers int, items []string, handler func(ctx context.Context, item string) error) error {
	var ptrs []*string
	for _, s := range items {
		str := s
		ptrs = append(ptrs, &str)
	}
	return ForEach[string](ctx, workers, ptrs, func(ctx context.Context, item *string) error {
		return handler(ctx, *item)
	})
}

// ForEach executes a handler function in parallel for each item in the items list
// WARNING: handler is not threadsafe. You may need to lock non-threadsafe resources or use other techniques
func ForEach[T any](ctx context.Context, workers int, items []*T, handler func(ctx context.Context, item *T) error) error {
	grp, ctx := errgroup.WithContext(ctx)
	ch := make(chan *T)

	for i := 0; i < workers; i++ {
		grp.Go(func() error {
			for item := range ch {
				select {
				case <-ctx.Done():
				default:
					it := item
					if err := handler(ctx, it); err != nil {
						return err
					}
				}
			}
			return nil
		})
	}
	grp.Go(func() error {
		defer close(ch)
		for _, i := range items {
			select {
			case <-ctx.Done():
			default:
				ch <- i
			}
		}
		return nil
	})

	return grp.Wait()
}
