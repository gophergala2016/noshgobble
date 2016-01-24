package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
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

type Ingredient struct {
	qty  float64
	unit string
	name string
}
type IngredientData struct {
	ingredients []Ingredient
	data        map[int]float64
}

func getIngredientData(db *sql.DB, ingredients string) (data IngredientData, err error) {
	data.ingredients = make([]Ingredient, 0)
	data.data = make(map[int]float64)
	for _, ingredient := range strings.Split(ingredients, "\n") {
		parser := NewParser(strings.NewReader(ingredient))
		item, err := parser.Parse()
		if err != nil {
			return nil, err
		}
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
			foodData[id] += qty
		}
		append
		fmt.Println(len(foodData))
	}
	return
}
