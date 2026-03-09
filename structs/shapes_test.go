package structs

import (
	"testing"
)

func TestShapes(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	t.Run("calculate the area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{Width: 12.0, Height: 6.0}
		checkArea(t, rectangle, 72.0)
	})
	t.Run("calculate the area of a circle", func(t *testing.T) {
		circle := Circle{Radius: 10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}

func testAreaTableDriven(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got :=tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}
}
