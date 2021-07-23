package iteration

import "testing"

func TestIter(t *testing.T) {
	got := Iter("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}
