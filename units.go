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
	factor   float64
	name     string
	abbr     string
	synonyms []string
}

var unitMap = map[Unit]*UnitModel{
	GRAM:        &UnitModel{GRAM, 1.0, "gram", "g", []string{"grams"}},
	KILOGRAM:    &UnitModel{KILOGRAM, 1000.0, "kilogram", "kg", []string{"kilograms"}},
	MILLIGRAM:   &UnitModel{MILLIGRAM, 0.001, "milligram", "mg", []string{"milligrams"}},
	MICROGRAM:   &UnitModel{MICROGRAM, 0.000001, "microgram", "µg", []string{"micrograms"}},
	CUP:         &UnitModel{CUP, 250.0, "cup", "cup", []string{"cups", "c"}},
	LITER:       &UnitModel{LITER, 1000.0, "liter", "L", []string{"liters", "l", "litre", "litres"}},
	MILLILITER:  &UnitModel{MILLILITER, 1.0, "milliliter", "mL", []string{"milliliters", "ml"}},
	OUNCE:       &UnitModel{OUNCE, 28.3495, "ounce", "oz", []string{"ounces"}},
	PINT:        &UnitModel{PINT, 473.176, "pint", "pt", []string{"pints"}},
	POUND:       &UnitModel{POUND, 453.592, "pound", "lb", []string{"pounds"}},
	TABLESPOON:  &UnitModel{TABLESPOON, 15.0, "tablespoon", "tbsp", []string{"tablespoons", "T", "tb", "tbl"}},
	TEASPOON:    &UnitModel{TEASPOON, 5.0, "teaspoon", "tsp", []string{"teaspoons", "t", "tsp"}},
	IU:          &UnitModel{IU, 0.0, "IU", "IU", []string{"iu", "internation-units"}},
	KILOCALORIE: &UnitModel{KILOCALORIE, 0.0, "kilocalorie", "kcal", []string{"kilocalories"}},
	KILOJOULE:   &UnitModel{KILOJOULE, 0.0, "kilojoule", "kJ", []string{"kilojoules"}},
	COUNT:       &UnitModel{COUNT, 0.0, "<>", "<>", []string{}},
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
