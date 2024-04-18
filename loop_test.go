package asyncloop

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestLoop(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	var sum int32 = 0

	startedAt := time.Now()

	Loop(slice, func(i int, v int) {
		time.Sleep(time.Millisecond * 100)
		atomic.AddInt32(&sum, int32(v*2))
	})

	finishedAt := time.Now()

	expectedSum := int32(2 + 4 + 6 + 8 + 10)
	if sum != expectedSum {
		t.Errorf("Expected sum to be %d, but got %d", expectedSum, sum)
	}

	elapsed := finishedAt.Sub(startedAt)
	if elapsed < time.Millisecond*100 || elapsed > time.Millisecond*200 {
		t.Errorf(
			"Expected elapsed time to be between 100ms and 200ms, but got %v",
			elapsed,
		)
	}
}

func TestLoopN(t *testing.T) {
	var sum int32 = 0

	startedAt := time.Now()

	LoopN(5, func(i int) {
		time.Sleep(time.Millisecond * 100)
		atomic.AddInt32(&sum, int32(i*2))
	})

	finishedAt := time.Now()

	expectedSum := int32(0 + 2 + 4 + 6 + 8)
	if sum != expectedSum {
		t.Errorf("Expected sum to be %d, but got %d", expectedSum, sum)
	}

	elapsed := finishedAt.Sub(startedAt)
	if elapsed < time.Millisecond*100 || elapsed > time.Millisecond*200 {
		t.Errorf(
			"Expected elapsed time to be between 100ms and 200ms, but got %v",
			elapsed,
		)
	}
}
