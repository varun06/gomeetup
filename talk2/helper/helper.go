package main

import "testing"

func failure(t *testing.T) {
	// This call to helper here silences this function in error reports.
	t.Helper()
	t.Fatal("failure")
}

func TestSomething(t *testing.T) {
	failure(t)
}
