package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jack")
	want := "hello, Jack"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
