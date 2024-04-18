package asyncloop

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestBatch(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	expected := map[string]struct{}{
		"[1 2 3]": {},
		"[4 5 6]": {},
		"[7 8 9]": {},
		"[10]":    {},
	}

	mu := sync.Mutex{}
	actualBatches := [][]int{}
	it := func(index int, batch []int) {
		time.Sleep(time.Millisecond * 100)
		mu.Lock()
		actualBatches = append(actualBatches, batch)
		mu.Unlock()
	}

	startedAt := time.Now()
	Batch(slice, 3, it)
	finishedAt := time.Now()

	elapsed := finishedAt.Sub(startedAt)
	if elapsed < time.Millisecond*100 || elapsed > time.Millisecond*200 {
		t.Errorf(
			"Expected elapsed time to be between 100ms and 200ms, but got %v",
			elapsed,
		)
	}

	for _, batch := range actualBatches {
		actual := fmt.Sprintf("%v", batch)

		if _, ok := expected[actual]; !ok {
			t.Errorf("Expected one of %v, got %s", expected, actual)
		}

	}
}
