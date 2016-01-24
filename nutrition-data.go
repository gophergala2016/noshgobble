package main

import "database/sql"

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

func getNutrientData(db *sql.DB, item *FoodItem) {
	//rows, err := db.Query("SELECT foods.description, units, description, precision FROM nutrients ORDER BY common_order")

}
