package asyncloop

import (
	"context"
	"sync"
)

// Pool is used to perform bounded concurrency when iterating over the
// elements in a slice.
//
// The workers parameter specifies the size of the concurrency pool
// for iteration. For example, if a factor of 2 is given, then there
// will only even be 2 iterations running at once.
// 1 would effectively be a serial iteration.
//
// Bounded concurrency is useful in cases where the user may wish
// to perform concurrency but in a reduced rate, so as to avoid
// rate limits or running out of file descriptors.
func Pool[T any](slice []T, workers int, it func(int, T) bool) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(len(slice))

	ch := make(chan struct{}, workers)
	defer close(ch)

	for i, v := range slice {
		ch <- struct{}{}

		go func(i int, v T) {
			defer func() {
				<-ch
				wg.Done()
			}()

			select {
			case <-ctx.Done():
				return
			default:
				if !it(i, v) {
					cancel()
					return
				}
			}
		}(i, v)
	}

	wg.Wait()
}
