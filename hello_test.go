package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Dan")
		want := "Hello, Dan!"
		
		assertCorrectingMessage(t, got, want)
	})

	t.Run("say 'Hello, World!' when supplied with an empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"
		
		assertCorrectingMessage(t, got, want)
	})
}

func assertCorrectingMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got,  want)
	}
}