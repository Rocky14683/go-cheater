package test

import (
	"testing"
)

func TestAdd(t *testing.T) {

	Add := func(x, y int) (res int) {
		return x + y
	}

	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestMinus(t *testing.T) {

	Minus := func(x, y int) (res int) {
		return x - y
	}

	got := Minus(4, 6)
	want := -2

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
