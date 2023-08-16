package vault

import "testing"

func TestFetch(t *testing.T) {
	vault := Create("aws", "asmnk-fjgnsd-sjfdb")
	result := Fetch(vault, "aws")
	expected := "asmnk-fjgnsd-sjfdb"

	if result != expected {
		t.Errorf("Result: %s Expected: %s", result, expected)
	}
}
