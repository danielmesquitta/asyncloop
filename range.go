package asyncloop

import "sync"

// Range creates a function iterator to iterate between two given
// integer like values.
//
// The first argument is the starting value, which is included in the
// iteration. The second argument is the stop value, which is when the
// iteration is stopped. This value is not included. The it function
// is called for each value in the iteration.
func Range(start int, stop int, it func(int)) {
	RangeWithStep(start, stop, 1, it)
}

// RangeWithStep creates a function iterator to iterate between two values
// with a given step incrementor.
//
// The first value is always returned (provided the stop value is value for
// the step amount)
// The stop value is not included.
// The step value can be either either greater than or less than 0. If the
// step is 0 then no iteration will take place.
// The it function is called for each value in the iteration.
func RangeWithStep(
	start,
	stop,
	step int,
	it func(int),
) {
	if step == 0 {
		return
	}

	delta := stop - start
	steps := delta / step
	rem := delta % step
	if rem > 0 {
		steps += 1
	}

	wg := sync.WaitGroup{}
	for i := 0; i < steps; i++ {
		wg.Add(1)
		num := i*step + start
		go func(num int) {
			defer wg.Done()
			it(num)
		}(num)
	}

	wg.Wait()
}
