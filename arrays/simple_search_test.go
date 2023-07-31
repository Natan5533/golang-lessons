package arrays

import (
	"reflect"
	"testing"
)

func TestListSearch(t *testing.T) {
	list := []int{5, 3, 5, 76, 2, 34}
	number := 5

	result := ListSearch(list, number)
	expected := []int{5, 5}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result: %v - Expected: %v", result, expected)
	}
}
