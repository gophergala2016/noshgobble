package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
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
		var unitString string
		err = rows.Scan(&nutrient.id, &unitString, &nutrient.name, &nutrient.precision)
		nutrient.unit = getUnitModel(unitString).id
		checkErr(err)
		nutrientOrder = append(nutrientOrder, nutrient.id)
		nutrients[nutrient.id] = nutrient
	}
}

func findFood(db *sql.DB, foodTerms string) (id int, err error) {
	rows, err := db.Query("SELECT id FROM food_fts WHERE food_fts MATCH ? ORDER BY bm25(food_fts) LIMIT 20", foodTerms)
	checkErr(err)
	if rows.Next() {
		err = rows.Scan(&id)
		checkErr(err)
	} else {
		id, err = -1, errors.New("No food found for the terms: "+foodTerms)
	}
	return
}

func getFoodName(db *sql.DB, id int) (name string, err error) {
	rows, err := db.Query("SELECT description FROM foods WHERE id = ?", id)
	if err != nil {
		return
	}
	if rows.Next() {
		err = rows.Scan(&name)
	} else {
		err = errors.New(fmt.Sprintf("No food found with the id: %d", id))
	}
	return
}

func getFoodData(db *sql.DB, foodId int) map[int]float64 {
	qtys := make(map[int]float64)
	rows, err := db.Query("SELECT nutrient_id, quantity FROM quantities WHERE food_id = ?", foodId)
	checkErr(err)
	for rows.Next() {
		var id int
		var qty float64
		err = rows.Scan(&id, &qty)
		checkErr(err)
		qtys[id] = qty
	}
	return qtys
}

func getIngredientData(db *sql.DB, ingredientStr string) (map[string]interface{}, error) {
	ingredients := make([][]interface{}, 0)
	data := make(map[int]float64)
	log.Printf("parsing `%s`", ingredientStr)
	for _, ingredient := range strings.Split(ingredientStr, "\n") {
		log.Printf("  > parsing `%s`", ingredient)
		parser := NewParser(strings.NewReader(ingredient))
		item, err := parser.Parse()
		if err != nil {
			if ingredient != "" {
				log.Printf("Failed to parse `%s`", ingredient)
			}
			continue
		}
		unit := unitMap[item.unit]
		log.Println(item.terms)
		foodId, err := findFood(db, item.terms)
		if err != nil {
			return nil, err
		}
		foodName, err := getFoodName(db, foodId)
		if err != nil {
			return nil, err
		}
		for id, qty := range getFoodData(db, foodId) {
			data[id] += qty * (item.quantity * unit.factor / 100)
		}

		ingredients = append(ingredients, []interface{}{item.quantity, unit.abbr, foodName})
	}
	dataMap := make(map[string]float64)
	for id, qty := range data {
		dataMap[strconv.Itoa(id)] = qty
	}

	return map[string]interface{}{"ingredients": ingredients, "data": dataMap}, nil
}
