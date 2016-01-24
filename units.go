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
	MILLIGRAM
	MICROGRAM
	IU
	KILOCALORIE
	KILOJOULE
	OUNCE
	PINT
	POUND
	TABLESPOON
	TEASPOON
	COUNT
)

type UnitModel struct {
	id       Unit
	name     string
	abbr     string
	synonyms []string
}

var unitMap = map[Unit]*UnitModel{
	GRAM:        &UnitModel{GRAM, "gram", "g", []string{"grams"}},
	MILLIGRAM:   &UnitModel{MILLIGRAM, "milligram", "mg", []string{"milligrams"}},
	MICROGRAM:   &UnitModel{MICROGRAM, "microgram", "µg", []string{"micrograms"}},
	CUP:         &UnitModel{CUP, "cup", "cup", []string{"cups", "c"}},
	KILOGRAM:    &UnitModel{KILOGRAM, "kilogram", "kg", []string{"kilograms"}},
	LITER:       &UnitModel{LITER, "liter", "L", []string{"liters", "l", "litre", "litres"}},
	MILLILITER:  &UnitModel{MILLILITER, "milliliter", "mL", []string{"milliliters", "ml"}},
	OUNCE:       &UnitModel{OUNCE, "ounce", "oz", []string{"ounces"}},
	PINT:        &UnitModel{PINT, "pint", "pt", []string{"pints"}},
	POUND:       &UnitModel{POUND, "pound", "lb", []string{"pounds"}},
	TABLESPOON:  &UnitModel{TABLESPOON, "tablespoon", "tbsp", []string{"tablespoons", "T", "tb", "tbl"}},
	TEASPOON:    &UnitModel{TEASPOON, "teaspoon", "tsp", []string{"teaspoons", "t", "tsp"}},
	IU:          &UnitModel{IU, "IU", "IU", []string{"iu", "internation-units"}},
	KILOCALORIE: &UnitModel{KILOCALORIE, "kilocalorie", "kcal", []string{"kilocalories"}},
	KILOJOULE:   &UnitModel{KILOJOULE, "kilojoule", "kJ", []string{"kilojoules"}},
	COUNT:       &UnitModel{COUNT, "<>", "<>", []string{}},
}

func (u Unit) String() string {
	return unitMap[u].abbr
}

var synonym2Model map[string]*UnitModel

func loadSynonym2Model() {
	synonym2Model = make(map[string]*UnitModel)

	for _, model := range unitMap {
		for _, synonym := range model.synonyms {
			synonym2Model[synonym] = model
		}
		synonym2Model[model.name] = model
		synonym2Model[model.abbr] = model
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
		return model.name
	} else {
		return ""
	}
}

func getUnitAbbr(w string) string {
	model := getUnitModel(w)
	if model != nil {
		return model.abbr
	} else {
		return ""
	}
}
