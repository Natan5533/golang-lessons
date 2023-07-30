package search

import "testing"

func TestSearch(t *testing.T) {
	list := [5]int{2, 3, 4, 1, 1}

	response := Search(list, 1)
	expected := 2

	if response != expected {
		t.Errorf("Result: %d, espected: %d", response, expected)
	}
}
