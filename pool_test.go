package asyncloop

import (
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	mu := sync.Mutex{}
	out := []int{}

	startedAt := time.Now()
	Pool(slice, 3, func(i int, v int) bool {
		time.Sleep(time.Millisecond * 50)
		mu.Lock()
		out = append(out, v)
		mu.Unlock()
		return i < 3
	})
	finishedAt := time.Now()

	elapsed := finishedAt.Sub(startedAt)
	if elapsed < time.Millisecond*100 || elapsed > time.Millisecond*200 {
		t.Errorf(
			"Expected elapsed time to be between 100ms and 200ms, but got %v",
			elapsed,
		)
	}

	expected := 6
	if len(out) != expected {
		t.Errorf("Expected %d elements, got %d", expected, len(out))
	}
}
