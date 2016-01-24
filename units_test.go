package main

import "testing"

func TestBaseUnit(t *testing.T) {
	for i, x := range []struct {
		arg      string
		expected string
	}{
		{"grans", ""},
		{"grams", "gram"},
		{"g", "gram"},
		{"litres", "liter"},
	} {
		if BaseUnit(x.arg) != x.expected {
			t.Errorf("Test %d: expected `%s` but saw `%s`\n", i, x.expected, BaseUnit(x.arg))
		}
	}
}
