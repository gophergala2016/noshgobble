package main

import (
	"database/sql"
	"log"
)

const dbFilename = "food.db"
const dbPath = "./databases/"
const foodTermsPath = dbPath + "food-terms.txt"

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func dbRun(db *sql.DB, sql string) {
	_, err := db.Exec(string(sql))
	checkErr(err)
}
