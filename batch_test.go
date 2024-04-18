package asyncloop

import (
	"fmt"
	"testing"
	"time"
)

func TestBatch(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	expected := map[string]int{
		"[1 2 3]": 0,
		"[4 5 6]": 1,
		"[7 8 9]": 2,
		"[10]":    3,
	}

	actualBatches := make([][]int, 4)
	it := func(index int, batch []int) {
		time.Sleep(time.Millisecond * 100)
		actualBatches[index] = batch
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

	for i, batch := range actualBatches {
		actual := fmt.Sprintf("%v", batch)

		j, ok := expected[actual]
		if !ok {
			t.Errorf("Expected one of %v, got %s", expected, actual)
		}

		if i != j {
			t.Errorf("Expected index %d, got %d", j, i)
		}
	}
}
