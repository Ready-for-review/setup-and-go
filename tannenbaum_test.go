package main

import (
	"testing"
)

func Test_PrintTannenbaum(t *testing.T) {
	got := PrintTannenbaum(1)
	want := "*"

	if got != want {
		t.Errorf("got %v but wanted %v", got, want)
	}
}
