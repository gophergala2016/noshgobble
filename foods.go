package main

import (
	"bufio"
	"os"

	"github.com/a2800276/porter"
)

var foodTerms map[string]bool

func loadFoodTerms() {
	foodTerms = make(map[string]bool)
	f, _ := os.Open(foodTermsPath)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		foodTerms[scanner.Text()] = true
	}
}

func IsFoodTerm(w string) bool {
	if foodTerms == nil {
		loadFoodTerms()
	}

	return foodTerms[porter.Stem(w)]
}
