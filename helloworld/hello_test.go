package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Yamuna")
	want := "Hello, Yamuna"

	if got != want {
		t.Errorf(" got %q want %q", got, want)
	}
}