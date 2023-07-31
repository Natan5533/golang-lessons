package arrays

import "testing"

func TestListLength(t *testing.T) {
	assert := func(t *testing.T, result, expected int, array []int) {
		t.Helper()
		if expected != result {
			t.Errorf("Result: %d - Expected: %d List: %v", result, expected, array)
		}
	}
	t.Run("slice", func(t *testing.T) {
		slice := []int{5, 5}

		result := ListCounter(slice)
		expected := 10

		assert(t, result, expected, slice)

	})
}
