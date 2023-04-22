package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying Hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, world' when an empty string is passed", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)

	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Nancy", "Spanish")
		want := "Hola, Nancy"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
