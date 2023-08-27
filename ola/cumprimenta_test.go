package ola

import (
	"bytes"
	"testing"
)

func TestCumprimenta(t *testing.T) {
	buffer := bytes.Buffer{}
	Cumprimenta(&buffer, "Natana")

	result := buffer.String()
	expected := "Olá Natana"

	if result != expected {
		t.Errorf("R: '%s', E: '%s'", result, expected)
	}

}
