package interacao

import (
	"fmt"
	"testing"
)

// Example
func ExampleRepeat() {
	response := Repeat("Monkey D. Luffy", 2)
	fmt.Println(response)
	// Output: Monkey D. LuffyMonkey D. Luffy
}

// Test
func TestRepeat(t *testing.T) {
	response := Repeat("a", 6)
	expected_response := "aaaaaa"

	if response != expected_response {
		t.Errorf("[Fail] Expected: '%s' Returned: '%s'", expected_response, response)
	}
}

// Benchmark
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 2)
	}
}
