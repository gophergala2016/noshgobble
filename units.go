package main

import "strings"

type Unit int

const (
	// Special tokens
	GRAM Unit = iota
	CUP
	KILOGRAM
	LITER
	MILLILITER
	OUNCE
	PINT
	POUND
	TABLESPOON
	TEASPOON
)

type UnitModel struct {
	id          Unit
	defaultName string
	synonyms    []string
}

var unitMap = map[Unit]UnitModel{
	GRAM:       {GRAM, "gram", []string{"grams", "g"}},
	CUP:        {CUP, "cup", []string{"cups", "c"}},
	KILOGRAM:   {KILOGRAM, "kilogram", []string{"kilograms", "kg"}},
	LITER:      {LITER, "liter", []string{"liters", "l", "litre", "litres"}},
	MILLILITER: {MILLILITER, "milliliter", []string{"milliliters", "ml"}},
	OUNCE:      {OUNCE, "ounce", []string{"ounces", "oz"}},
	PINT:       {PINT, "pint", []string{"pints", "pt"}},
	POUND:      {POUND, "pound", []string{"pounds", "lb"}},
	TABLESPOON: {TABLESPOON, "tablespoon", []string{"tablespoons", "T", "tb", "tbl", "tbsp"}},
	TEASPOON:   {TEASPOON, "teaspoon", []string{"teaspoons", "t", "tsp", "tbl", "tbsp"}},
}

var synonym2Model map[string]UnitModel

func loadSynonym2Model() {
	synonym2Model = make(map[string]UnitModel)

	for _, model := range unitMap {
		for _, synonym := range model.synonyms {
			synonym2Model[synonym] = model
		}
		synonym2Model[model.defaultName] = model
	}
}

func BaseUnit(w string) string {
	if synonym2Model == nil {
		loadSynonym2Model()
	}

	model, ok := synonym2Model[w]
	if !ok {
		model, ok = synonym2Model[strings.ToLower(w)]
	}
	if !ok {
		return ""
	} else {
		return model.defaultName
	}
}
