package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("jack")
		want := "hello, jack"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say hello empty", func(t *testing.T) {
		got := Hello("")
		want := "hello, world"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
