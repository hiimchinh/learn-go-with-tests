package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chinh")
	want := "Hello, Chinh"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
