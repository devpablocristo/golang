package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello world!ddd"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestMultiply(t *testing.T) {
	t.Run("multiply 2 * 3", func(t *testing.T) {
		got := Multiply(2, 3)
		want := 6

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("multiply by 3 * 5", func(t *testing.T) {
		got := Multiply(3, 5)
		want := 15

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("multiply by 4 * 4", func(t *testing.T) {
		got := Multiply(3, 5)
		want := 198

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestHello2(t *testing.T) {
	t.Run("s is empty", func(t *testing.T) {
		got := Hello2("")
		want := "Hello world!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("s is not empty", func(t *testing.T) {
		got := Hello2("Tori")
		want := "Hello Tori!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
