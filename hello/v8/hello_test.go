package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("to people, in French", func(t *testing.T) {
		got := Hello("Chris", "French")
		want := "Bonjour, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("to people, in Spanish", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty name, with non-empty language", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, mundo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty name, with non-empty language", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, monde"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty name and empty language, defaults to 'Hello world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	fmt.Printf("got %q want %q\n", got, want)
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
