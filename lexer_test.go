package main

import (
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	type token struct {
		tok Token
		lit string
	}
	for i, x := range []struct {
		input  string
		output []token
	}{
		{"2t	of   sugar", []token{
			{NUMBER, "2"},
			{UNIT, "teaspoon"},
			{WHITESPACE, " "},
			{OF, "of"},
			{WHITESPACE, " "},
			{FOOD, "sugar"},
		}},
		{"2 T of curry powder", []token{
			{NUMBER, "2"},
			{WHITESPACE, " "},
			{UNIT, "tablespoon"},
			{WHITESPACE, " "},
			{OF, "of"},
			{WHITESPACE, " "},
			{FOOD, "curry"},
			{WHITESPACE, " "},
			{FOOD, "powder"},
		}},
	} {
		scanner := NewScanner(strings.NewReader(x.input))
		for j, expected := range x.output {
			tok, lit := scanner.Scan()
			if tok != expected.tok || lit != expected.lit {
				t.Errorf("Test %d, token %d: expected %s:`%s` but got %s:`%s`",
					i, j, expected.tok, expected.lit, tok, lit)
			}
		}
		tok, lit := scanner.Scan()
		if tok != EOF {
			t.Errorf("Test %d: expected EOF but got %d:%s", i, tok, lit)
		}
	}
}
