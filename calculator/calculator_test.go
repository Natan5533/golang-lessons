package calculator

import "testing"

func TestCalculator(t *testing.T) {
	verifyResponse := func(t *testing.T, response, expected int) {
		t.Helper()
		if response != expected {
			t.Errorf("[Fail] Expected: '%d' Returned '%d'", expected, response)
		}
	}
	t.Run("test mult", func(t *testing.T) {
		response := Mult(5, 2)
		expected_response := 10

		verifyResponse(t, response, expected_response)
	})
	t.Run("test sum", func(t *testing.T) {
		response := Sum(5, 5)
		expected_response := 10

		verifyResponse(t, response, expected_response)
	})
	t.Run("test div", func(t *testing.T){
		response := Div(10, 2)
		expected_response := 5

		verifyResponse(t, response, expected_response)
	})
	t.Run("test sub", func(t *testing.T){
		response := Sub(10, 5)
		expected_response := 5

		verifyResponse(t, response, expected_response)
	})
}
