package main

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("v")
	exp := "vvvvv"

	if got != exp {
		t.Errorf("expected %q got %q", exp, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}
