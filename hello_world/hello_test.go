package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("should return name if passed in", func(t *testing.T) {
		got := Hello("Shensen")
		want := "Hello, Shensen"
		assertCorrectMessage(t, got, want)
	})

	t.Run("should respond with default if empty string supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
