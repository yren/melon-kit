package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestHello(t *testing.T) {
	assertMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say Chinese hello", func(t *testing.T) {
		got := Hello("jack", "Chinese")
		want := "你好, jack"
		assertMessage(t, got, want)
	})

	t.Run("say Spanish hello", func(t *testing.T) {
		got := Hello("Mary", "Spanish")
		want := "Hola, Mary"
		assertMessage(t, got, want)
	})

	t.Run("default say English", func(t *testing.T) {
		got := Hello("", "")
		want := "hello, world"
		assertMessage(t, got, want)
	})
}

func TestSplit(t *testing.T) {
	latlng := ""
	splitted := strings.Split(latlng, ",")
	if len(splitted) != 2 {
		fmt.Println(len(splitted))
	}
}
