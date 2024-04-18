package asyncloop

import "sync"

// Batch is used to turn any slice into an iterator of batches, with the size
// of each batch being the second parameter.
//
// This is useful if you want to perform batch operations on a large slice
// of elements, for example, breaking up a large request into multiple
// smaller ones.
func Batch[T any](slice []T, size uint, it func(int, []T) bool) {
	if size < 1 {
		return
	}

	index := 0
	wg := sync.WaitGroup{}

	for i := uint(0); i < uint(len(slice)); i += size {
		wg.Add(1)

		top := min(uint(len(slice)), uint(i)+size)
		batch := slice[i:top]

		go func() {
			defer wg.Done()
			it(index, batch)
		}()

		index++
	}

	wg.Wait()
}
