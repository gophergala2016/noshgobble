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
	COUNT
)

type UnitModel struct {
	id          Unit
	defaultName string
	synonyms    []string
}

var unitMap = map[Unit]*UnitModel{
	GRAM:       &UnitModel{GRAM, "gram", []string{"grams", "g"}},
	CUP:        &UnitModel{CUP, "cup", []string{"cups", "c"}},
	KILOGRAM:   &UnitModel{KILOGRAM, "kilogram", []string{"kilograms", "kg"}},
	LITER:      &UnitModel{LITER, "liter", []string{"liters", "l", "litre", "litres"}},
	MILLILITER: &UnitModel{MILLILITER, "milliliter", []string{"milliliters", "ml"}},
	OUNCE:      &UnitModel{OUNCE, "ounce", []string{"ounces", "oz"}},
	PINT:       &UnitModel{PINT, "pint", []string{"pints", "pt"}},
	POUND:      &UnitModel{POUND, "pound", []string{"pounds", "lb"}},
	TABLESPOON: &UnitModel{TABLESPOON, "tablespoon", []string{"tablespoons", "T", "tb", "tbl", "tbsp"}},
	TEASPOON:   &UnitModel{TEASPOON, "teaspoon", []string{"teaspoons", "t", "tsp", "tbl", "tbsp"}},
	COUNT:      &UnitModel{COUNT, "<>", []string{}},
}

func (u Unit) String() string {
	return unitMap[u].defaultName
}

var synonym2Model map[string]*UnitModel

func loadSynonym2Model() {
	synonym2Model = make(map[string]*UnitModel)

	for _, model := range unitMap {
		for _, synonym := range model.synonyms {
			synonym2Model[synonym] = model
		}
		synonym2Model[model.defaultName] = model
	}
}

func getUnitModel(w string) *UnitModel {
	if synonym2Model == nil {
		loadSynonym2Model()
	}

	model, ok := synonym2Model[w]
	if !ok {
		model, ok = synonym2Model[strings.ToLower(w)]
	}
	if !ok {
		return nil
	} else {
		return model
	}
}

func getUnitName(w string) string {
	model := getUnitModel(w)
	if model != nil {
		return model.defaultName
	} else {
		return ""
	}
}
