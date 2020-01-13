package _5_structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}

// Table Driven Tests

func TestAreaTableDriven(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}

}

func TestPerimeterTableDriven(t *testing.T) {

	perimeterTests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasPerimeter: 36.0},
		{name: "Circle", shape: Circle{10}, hasPerimeter: 62.83185307179586},
		{name: "Triangle", shape: Triangle{12, 6}, hasPerimeter: 18.0},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.hasPerimeter {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasPerimeter)
		}
	}
}
