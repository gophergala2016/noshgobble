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

func TestGetFoodName(t *testing.T) {
	// open database
	db, err := sql.Open("sqlite3", dbPath+dbFilename)
	checkErr(err)
	defer func() {
		err := db.Close()
		checkErr(err)
	}()

	for i, x := range []struct {
		foodId   int
		foodName string
	}{
		{1000, ""},
		{1001, "Butter, salted"},
		{1002, "Butter, whipped, with salt"},
		{1003, "Butter oil, anhydrous"},
	} {
		foodName, err := getFoodName(db, x.foodId)
		if x.foodName == "" && err == nil {
			t.Errorf("Test %d: error was expected", i)
		} else if x.foodName != "" && err != nil {
			t.Errorf("Test %d: Error: %v", i, err)
		}
		if foodName != x.foodName {
			t.Errorf("Test %d: food names differ - expected %s but got %s", i, x.foodName, foodName)
		}
	}
}

func TestGetFoodData(t *testing.T) {
	db, err := sql.Open("sqlite3", dbPath+dbFilename)
	checkErr(err)
	defer func() {
		err := db.Close()
		checkErr(err)
	}()

	loadNutrients(db)
	nutQtys := getFoodData(db, 19362)
	//for _, id := range nutrientOrder {
	//	if qty, ok := nutQtys[id]; ok {
	//		fmt.Println(id, nutrients[id].name, qty)
	//	}
	//}
	if nutQtys[205] != 83.9 {
		t.Errorf("Found unexpected data in nutrient quantity database")
	}
}
