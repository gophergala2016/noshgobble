package main

import (
	"database/sql"
	"testing"
)

func TestFindFood(t *testing.T) {
	// open database
	db, err := sql.Open("sqlite3", dbPath+dbFilename)
	checkErr(err)
	defer func() {
		err := db.Close()
		checkErr(err)
	}()

	for i, x := range []struct {
		foodTerms string
		foodId    int
	}{
		{"sugar", 28210},
		{"water", 14555},
		{"baking soda", 18372},
		{"carrot", 11683},
	} {
		foodId, err := findFood(db, x.foodTerms)
		if err != nil {
			t.Errorf("Test %d: Error: %v", i, err)
		}
		if foodId != x.foodId {
			t.Errorf("Test %d: foodIds differ - expected %d but got %d", i, x.foodId, foodId)
		}
	}
}
