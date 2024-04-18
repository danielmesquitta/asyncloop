package asyncloop

import (
	"sync"
)

// Loop provides the ability to range over a slice concurrently.
// Each element of the slice will be called within it's own goroutine,
// passing the index and value to the function it.
//
// This function should not be used in hopes to speed up any pure compute
// operation as there is an associated cost with spawning a new goroutine.
// Instead, it makes sense if there are any long running tasks inside of
// your loop.
//
// The TestLoop in Loop_test.go show a good example of when this
// method will speed up performance. (using time.Sleep)
func Loop[T any](slice []T, it func(int, T)) {
	wg := sync.WaitGroup{}
	wg.Add(len(slice))
	for i, v := range slice {
		go func() {
			defer wg.Done()
			it(i, v)
		}()
	}
	wg.Wait()
}

// LoopN allows you to perform a concurrent iteration n times.
//
// This is very similar to the Loop function, except that instead
// of looping over a slice of elements, it instead will loop from 0 to n-1.
func LoopN(n int, it func(int)) {
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			it(i)
		}()
	}
	wg.Wait()
}
