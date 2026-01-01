package main

import "testing"

func Test_Hello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Shreyash", "")
		want := "Hello, Shreyash"
		assertCorrectMessage(t, got, want)

	})

	t.Run("say 'Hello World', when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Kyra", "Spanish")
		want := "Hola, Kyra"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Amelie", "French")
		want := "Bonjour, Amelie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Amelie", "French")
		want := "Bonjour, Amelie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in German", func(t *testing.T) {
		got := Hello("Hanz", "German")
		want := "Halo, Hanz"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
