package main

import "testing"

func TestHello(t *testing.T) {
	assertMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("jack")
		want := "hello, jack"
		assertMessage(t, got, want)
	})

	t.Run("empty hello", func(t *testing.T) {
		got := Hello("")
		want := "hello, world"
		assertMessage(t, got, want)
	})
}
