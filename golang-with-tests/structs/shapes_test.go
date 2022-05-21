package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{5.8, 8.9}
	result := rectangle.Perimeter()
	expected := 29.4

	if result != expected {
		t.Errorf("Result: '%f' Expected: '%f'", result, expected)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		result := shape.Area()

		if result != expected {
			t.Errorf("Result: '%f' Expected: '%f'", result, expected)
		}
	}

	t.Run("area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{26.0, 42.0}

		checkArea(t, rectangle, 1092.0)
	})

	t.Run("area of a circle", func(t *testing.T) {
		circle := Circle{10}

		checkArea(t, circle, 314.1592653589793)
	})
}
