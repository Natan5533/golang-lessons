package v1

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Increment counter 3 times.", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assert(t, counter, 3)
	})
	t.Run("runs safely concurrently", func(t *testing.T) {
		expected := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(expected)

		for i := 0; i < expected; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assert(t, counter, expected)
	})
}

func assert(t *testing.T, got *Counter, expected int) {
	t.Helper()

	result := got.Show()

	if result != expected {
		t.Errorf("Expected: %d Got: %d", expected, result)
	}
}
