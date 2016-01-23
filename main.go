package main

import (
	"database/sql"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"time"

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
		t := time.Now()
		log.Print("Loading data into database!!", t.Format(time.RFC3339))
		// copy database to backup
		if err := os.Mkdir("./backups", 0755); err != nil && err.Error() != "mkdir ./backups: file exists" {
			checkErr(err)
		}
		sqlite3db, err := ioutil.ReadFile("./sqlite3.db")
		checkErr(err)
		err = ioutil.WriteFile("./backups/sqlite3.db"+t.Format(time.RFC3339), sqlite3db, 0640)
		checkErr(err)

		err = os.Remove("./sqlite3.db")
		checkErr(err)

		db, err := sql.Open("sqlite3", "./sqlite3.db")
		checkErr(err)
		sql, err := ioutil.ReadFile("./sql/create-food-table.sql")
		checkErr(err)
		_, err = db.Exec(string(sql))
		checkErr(err)
	}
}
