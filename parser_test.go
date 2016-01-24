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
		{"2t	of   sugar", FoodItem{2.0, TEASPOON, "sugar"}},
		{"2.6 T of curry powder", FoodItem{2.6, TABLESPOON, "curry powder"}},
		{"TWO THIRDS OF A CUP OF WATER", FoodItem{0.666, CUP, "water"}},
		{"3 grams of cheese\n", FoodItem{3, GRAM, "cheese"}},
	} {
		parser := NewParser(strings.NewReader(x.input))
		item, err := parser.Parse()
		if err != nil {
			t.Errorf("Test %d: Error: %v", i, err)
		}
		if item.quantity != x.item.quantity {
			t.Errorf("Test %d: quantities differ - expected %f but got %f", i, x.item.quantity, item.quantity)
		}
		if item.unit != x.item.unit {
			t.Errorf("Test %d: units differ - expected %s but got %s", i, x.item.unit, item.unit)
		}
		if item.terms != x.item.terms {
			t.Errorf("Test %d: terms differ - expected `%s` but got `%s`", i, x.item.terms, item.terms)
		}
	}
}
