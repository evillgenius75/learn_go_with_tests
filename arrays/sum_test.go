package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("wanted %d, but got %d from %v", want, got, numbers)
		}
	})
}

func TestSummAll(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %v but got %v", want, got)
		}
	}

	t.Run("variadic sum function with multiple slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		checkSums(t, got, want)

	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %v but got %v", want, got)
		}
	}

	t.Run("variadic sum function with multiple slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely run if an empty slice is given", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1}, []int{0, 9})
		want := []int{0, 0, 9}

		checkSums(t, got, want)
	})
}
