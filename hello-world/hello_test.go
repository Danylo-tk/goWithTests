package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Dan", "")
		want := "Hello, Dan!"
		assertCorrectingMessage(t, got, want)
	})

	t.Run("say 'Hello, World!' when supplied with an empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectingMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Diego", "Spanish")
		want := "Hola, Diego!"
		assertCorrectingMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Louis", "French")
		want := "Bonjour, Louis!"
		assertCorrectingMessage(t, got, want)
	})

	t.Run("in Ukrainian", func(t *testing.T) {
		got := Hello("Данило", "Ukrainian")
		want := "Здорові були, Данило!"
		assertCorrectingMessage(t, got, want)
	})
}

func assertCorrectingMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got,  want)
	}
}