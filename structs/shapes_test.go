package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("Perimeter tests:", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		want := 40.0
		if got := Perimeter(rectangle); got != want {
			t.Errorf("Perimeter() = %v, want %v", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("Area tests:", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		want := 72.0
		if got := Area(rectangle); got != want {
			t.Errorf("Area() = %v, want %v", got, want)
		}
	})
}
