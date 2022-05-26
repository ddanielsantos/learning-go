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
	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{
			name: "Rectangle",
			shape: Rectangle{
				height: 26.0,
				width:  42.0,
			},
			expected: 1092.0,
		},
		{
			name: "Circle",
			shape: Circle{
				radius: 10,
			},
			expected: 314.1592653589793,
		},
		{
			name: "Triangle",
			shape: Triangle{
				base: 12, height: 6,
			},
			expected: 36.0,
		},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.shape.Area()
			if result != tt.expected {
				t.Errorf("Area of %#v: '%f' Expected: '%f'", tt.shape, result, tt.expected)
			}
		})
	}
}
