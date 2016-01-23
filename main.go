package main

import (
	"database/sql"
	"encoding/csv"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type quoteFixReader struct {
	r        *os.File
	buffer   []byte
	index    int
	capacity int
	done     bool
}

func newQuoteFixReader(r *os.File) (qfr *quoteFixReader) {
	qfr = new(quoteFixReader)
	qfr.r = r
	qfr.buffer = make([]byte, 1000)
	qfr.index = 0
	qfr.capacity = 0
	qfr.done = false
	return
}

func (r *quoteFixReader) Read(b []byte) (n int, err error) {
	i := 0
	for ; i < len(b); i++ {
		if r.index == r.capacity {
			r.index = 0
			if r.done {
				return 0, io.EOF
			} else {
				r.capacity, err = (*(r.r)).Read(r.buffer)
				if err == io.EOF {
					if i == 0 {
						return 0, io.EOF
					} else {
						r.done = true
						return i, nil
					}
				} else {
					checkErr(err)
				}
				//r.buffer = bytes.ToLower(r.buffer)
			}
		}
		c := r.buffer[r.index]
		if c == '~' {
			b[i] = '"'
		} else if c == '"' {
			if i < len(b)-1 {
				b[i] = '"'
				i++
				b[i] = '"'
			} else {
				return i, nil
			}
		} else {
			b[i] = c
		}
		r.index++
	}
	return i, nil
}

var fLoad = flag.Bool("load", false, "Load the data from the original data files")

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func loadFoods(db *sql.DB) {
	sql, err := ioutil.ReadFile("./sql/create-foods-table.sql")
	checkErr(err)
	_, err = db.Exec(string(sql))
	checkErr(err)

	// prepare the insert statement
	stmt, err := db.Prepare("INSERT INTO foods(id, food_group_id, description, short_description, common_name, manufacturer_name, refuse_description, refuse, scientific_name, nitrogen_factor, protein_factor, fat_factor, carbohydrate_factor) values(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	f, err := os.Open("./data/FOOD_DES.txt")
	checkErr(err)
	csvReader := csv.NewReader(newQuoteFixReader(f))
	csvReader.Comma = '^'
	for {
		l, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		checkErr(err)

		// insert
		_, err = stmt.Exec(l[0], l[1], l[2], l[3], l[4], l[5], l[7], l[8], l[9], l[10], l[11], l[12], l[13])
		checkErr(err)
	}
}

func loadNutrients(db *sql.DB) {
	sql, err := ioutil.ReadFile("./sql/create-nutrients-table.sql")
	checkErr(err)
	_, err = db.Exec(string(sql))
	checkErr(err)

	// prepare the insert statement
	stmt, err := db.Prepare("INSERT INTO nutrients(id, units, tagname, description, precision, common_order) values(?,?,?,?,?,?)")
	checkErr(err)

	f, err := os.Open("./data/NUTR_DEF.txt")
	checkErr(err)
	csvReader := csv.NewReader(newQuoteFixReader(f))
	csvReader.Comma = '^'
	for {
		l, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		checkErr(err)

		// insert
		_, err = stmt.Exec(l[0], l[1], l[2], l[3], l[4], l[5])
		checkErr(err)
	}
}

func loadQuantities(db *sql.DB) {
	sql, err := ioutil.ReadFile("./sql/create-quantities-table.sql")
	checkErr(err)
	_, err = db.Exec(string(sql))
	checkErr(err)

	// prepare the insert statement
	stmt, err := db.Prepare("INSERT INTO quantities(food_id, nutrient_id, quantity) values(?,?,?)")
	checkErr(err)

	f, err := os.Open("./data/NUT_DATA.txt")
	checkErr(err)
	csvReader := csv.NewReader(newQuoteFixReader(f))
	csvReader.Comma = '^'
	for {
		l, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		checkErr(err)

		// insert
		_, err = stmt.Exec(l[0], l[1], l[2])
		checkErr(err)
	}
}

func backupDB() {
	t := time.Now()
	if err := os.Mkdir("./backups", 0755); err != nil && err.Error() != "mkdir ./backups: file exists" {
		checkErr(err)
	}

	// copy database to backup
	sqlite3Bytes, err := ioutil.ReadFile("./sqlite3.db")
	checkErr(err)
	err = ioutil.WriteFile("./backups/sqlite3.db"+t.Format(time.RFC3339), sqlite3Bytes, 0640)
	checkErr(err)

	// remove old db file
	err = os.Remove("./sqlite3.db")
	checkErr(err)
}

func loadData() {
	log.Print("Loading data into database!!")

	backupDB()

	db, err := sql.Open("sqlite3", "./sqlite3.db")
	checkErr(err)
	defer func() {
		err := db.Close()
		checkErr(err)
	}()

	_, err = db.Exec("BEGIN TRANSACTION")
	checkErr(err)
	defer func() {
		_, err := db.Exec("END TRANSACTION")
		checkErr(err)
	}()

	// open database for writing
	loadFoods(db)
	loadNutrients(db)
	loadQuantities(db)
}

func main() {
	flag.Parse()
	if *fLoad {
		loadData()
	}
}
