package main

import "strings"

var unit2Synonym = map[string][]string{
	"gram":       {"grams", "g"},
	"liter":      {"liters", "l", "litre", "litres"},
	"tablespoon": {"tablespoons", "T", "tb", "tbl", "tbsp"},
	"teaspoon":   {"teaspoons", "t", "tsp", "tbl", "tbsp"},
	"cup":        {"cups", "c"},
	"kilogram":   {"kilograms", "kg"},
	"pound":      {"pounds", "lb"},
	"milliliter": {"milliliters", "ml"},
	"ounce":      {"ounces", "oz"},
	"pint":       {"pints", "pt"},
}

var synonym2Unit map[string]string

func loadSynonym2Unit() {
	synonym2Unit = make(map[string]string)

	for unit, synonyms := range unit2Synonym {
		for _, synonym := range synonyms {
			synonym2Unit[synonym] = unit
		}
	}
}

func BaseUnit(w string) string {
	if synonym2Unit == nil {
		loadSynonym2Unit()
	}

	unit := synonym2Unit[w]
	if unit == "" {
		return synonym2Unit[strings.ToLower(w)]
	} else {
		return unit
	}
}
