package main

import "testing"

func TestHello(t *testing.T) {
	verifyCorrectString := func(t *testing.T, result, expected string) { // functions are first class citizens in Go
		t.Helper() // especify that the t.Helper() is a helper function for the current test
		if result != expected {
			t.Errorf("Result: '%s' Expected: '%s'", result, expected)
		}
	}

	t.Run("say hello to a pearson", func(t *testing.T) {
		result := Hello("Chris")
		expected := "Hello, Chris!"

		verifyCorrectString(t, result, expected)
	})

	t.Run("say Hello, World when a empty string is supplied", func(t *testing.T) {
		result := Hello("")
		expected := "Hello, World!"

		verifyCorrectString(t, result, expected)
	})
}
