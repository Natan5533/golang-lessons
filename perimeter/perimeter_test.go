package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	result := Perimeter(rectangle)
	expected := 40.0

	if result != expected {
		t.Errorf("Result: %.2f Expected: %.2f", result, expected)
	}
}

func TestArea(t *testing.T) {
	verifyArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		result := shape.Area()

		if result != expected {
			t.Errorf("Result: %.2f Expected: %.2f", result, expected)
		}
	}
	t.Run("Rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		expected := 72.0

		verifyArea(t, rectangle, expected)
	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10}
		expected := 314.1592653589793

		verifyArea(t, circle, expected)
	})
}

// table driven test example

func TestAnyArea(t *testing.T) {
	expectedCircleArea := 314.1592653589793
	expectedRectangleArea := 60.00
	expectderiangleArea := 36.0

	anyArea := []struct {
		shape    Shape
		expected float64
	}{
		{Rectangle{Width: 12, Height: 5}, expectedRectangleArea},
		{Circle{Ray: 10}, expectedCircleArea},
		{Triangle{Base: 12, Height: 6}, expectderiangleArea},
	}

	for _, tt := range anyArea {
		result := tt.shape.Area()
		if result != tt.expected {
			t.Errorf("Failed! Result: %.13f, Expected %.13f", result, tt.expected)
		}
	}
}
