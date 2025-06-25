package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying Hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		exp := "Hello, Chris"

		assertCorrectMessage(t, got, exp)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		exp := "Hello, World"

		assertCorrectMessage(t, got, exp)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		exp := "Hola, Elodie"
		assertCorrectMessage(t, got, exp)
	})
}

func assertCorrectMessage(t testing.TB, got, exp string) {
	t.Helper()
	if got != exp {
		t.Errorf("got %q, expected %q", got, exp)
	}
}
