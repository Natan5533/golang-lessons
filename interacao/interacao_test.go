package interacao

import "testing"

func TestRepeat(t *testing.T) {
	response := Repeat("a", 6)
	expected_response := "aaaaaa"

	if response != expected_response {
		t.Errorf("[Fail] Expected: '%s' Returned: '%s'", expected_response, response)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 2)
	}
}
