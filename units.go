package main

var unit2Synonym = map[string][]string{
	"gram":  {"g", "grams"},
	"liter": {"l", "litre", "litres", "liters"},
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

	return synonym2Unit[w]
}
