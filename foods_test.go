package main

import "testing"

func TestIsFoodTerm(t *testing.T) {
	for i, x := range []struct {
		term     string
		expected bool
	}{
		{"zucchini", true},
		{"Tomatoes", true},
		{"POTATOES", true},
		{"letTuCes", true},
		{"koala", false},
		{"something", false},
		{"paper", false},
	} {
		if IsFoodTerm(x.term) != x.expected {
			t.Errorf("Test %d, food %s: expected %v but saw %v\n", i, x.term, x.expected, IsFoodTerm(x.term))
		}
	}
}
