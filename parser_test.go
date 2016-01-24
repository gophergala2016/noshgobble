package main

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	for i, x := range []struct {
		input string
		item  FoodItem
	}{
		{"2t	of   sugar", FoodItem{2, TEASPOON, "sugar"}},
		{"2 T of curry powder", FoodItem{2, TABLESPOON, "curry powder"}},
		{"TWO THIRDS OF A CUP OF WATER", FoodItem{0.666, CUP, "water"}},
	} {
		parser := NewParser(strings.NewReader(x.input))
		item, err := parser.Parse()
		if item.quantity != x.item.quantity {
			t.Errorf("Test %d: quantities differ - expected %f but got %f", i, item.quantity, x.item.quantity)
		}
		if item.unit != x.item.unit {
			t.Errorf("Test %d: units differ - expected %d but got %d", i, item.unit, x.item.unit)
		}
		if item.description != x.item.description {
			t.Errorf("Test %d: descriptions differ - expected `%s` but got `%s`", i, item.description, x.item.description)
		}
	}
}
