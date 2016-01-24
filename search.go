package main

import (
	"database/sql"
	"errors"
)

type FoodItem struct {
	quantity float64
	unit     Unit
	terms    string
}

type Nutrient struct {
	id        int
	unit      Unit
	name      string
	precision int
}

var nutrients map[int]Nutrient
var nutrientOrder []int

func loadNutrients(db *sql.DB) {
	nutrients = make(map[int]Nutrient)
	nutrientOrder = make([]int, 0, 150)

	db, err := sql.Open("sqlite3", dbPath+dbFilename)
	checkErr(err)
	defer func() {
		err := db.Close()
		checkErr(err)
	}()

	rows, err := db.Query("SELECT id, units, description, precision FROM nutrients ORDER BY common_order")
	checkErr(err)
	for rows.Next() {
		var nutrient Nutrient
		err = rows.Scan(&nutrient.id)
		checkErr(err)
		err = rows.Scan(&nutrient.unit)
		checkErr(err)
		err = rows.Scan(&nutrient.name)
		checkErr(err)
		err = rows.Scan(&nutrient.precision)
		checkErr(err)
		nutrientOrder = append(nutrientOrder, nutrient.id)
		nutrients[nutrient.id] = nutrient
	}
}

func findFood(db *sql.DB, foodTerms string) (int, error) {
	rows, err := db.Query("SELECT id FROM food_fts WHERE food_fts MATCH ? ORDER BY bm25(food_fts) LIMIT 20", foodTerms)
	checkErr(err)
	if rows.Next() {
		var id int
		err = rows.Scan(&id)
		checkErr(err)
		return id, nil
	} else {
		return -1, errors.New("No food found for the terms: " + foodTerms)
	}
}
