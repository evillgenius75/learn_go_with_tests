package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("Perimeter tests:", func(t *testing.T) {
		want := 40.0
		if got := Perimeter(10.0, 10.0); got != want {
			t.Errorf("Perimeter() = %v, want %v", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("Area tests:", func(t *testing.T) {
		want := 100.0
		if got := Area(10.0, 10.0); got != want {
			t.Errorf("Area() = %v, want %v", got, want)
		}
	})
}
