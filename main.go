package main

import (
	"database/sql"
	"flag"
	"io/ioutil"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var loadData = flag.Bool("load", false, "Load the data from the original data files")

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	if *loadData {
		log.Print("Loading data into database!!")
		db, err := sql.Open("sqlite3", "./sqlite3.db")
		checkErr(err)
		sql, err := ioutil.ReadFile("sql/create-food-table.sql")
		checkErr(err)
		_, err = db.Exec(string(sql))
		checkErr(err)
	}
}
