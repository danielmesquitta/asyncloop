package asyncloop

import (
	"sync"
	"testing"
	"time"
)

func TestRange(t *testing.T) {
	out := []int{}
	mu := sync.Mutex{}

	startedAt := time.Now()
	Range(0, 5, func(i int) {
		time.Sleep(time.Millisecond * 100)

		mu.Lock()
		out = append(out, i)
		mu.Unlock()
	})
	finishedAt := time.Now()

	elapsed := finishedAt.Sub(startedAt)
	if elapsed < time.Millisecond*100 || elapsed > time.Millisecond*200 {
		t.Errorf(
			"Expected elapsed time to be between 100ms and 200ms, but got %v",
			elapsed,
		)
	}

	expected := map[int]struct{}{0: {}, 1: {}, 2: {}, 3: {}, 4: {}}
	if len(out) != len(expected) {
		t.Errorf("Expected %d elements, got %d", len(expected), len(out))
	}

	for _, v := range out {
		if _, ok := expected[v]; !ok {
			t.Errorf("Expected one of %v, got %d", expected, v)
		}
	}
}

func TestRangeWithStep(t *testing.T) {
	out := []int{}
	mu := sync.Mutex{}

	startedAt := time.Now()
	RangeWithStep(0, 10, 2, func(i int) {
		time.Sleep(time.Millisecond * 100)

		mu.Lock()
		out = append(out, i)
		mu.Unlock()
	})
	finishedAt := time.Now()

	elapsed := finishedAt.Sub(startedAt)
	if elapsed < time.Millisecond*100 || elapsed > time.Millisecond*200 {
		t.Errorf(
			"Expected elapsed time to be between 100ms and 200ms, but got %v",
			elapsed,
		)
	}

	expected := map[int]struct{}{0: {}, 2: {}, 4: {}, 6: {}, 8: {}}
	if len(out) != len(expected) {
		t.Errorf("Expected %d elements, got %d", len(expected), len(out))
	}

	for _, v := range out {
		if _, ok := expected[v]; !ok {
			t.Errorf("Expected one of %v, got %d", expected, v)
		}
	}
}
