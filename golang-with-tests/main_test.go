package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Chris")
	expected := "Hello, Chris!"

	if result != expected {
		t.Errorf("Result: '%s' Expected: '%s'", result, expected)
	}
}
