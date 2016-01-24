package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"html/template"
	"log"
	"net/http"

	_ "github.com/dbalmain/go-sqlite3"
)

var fReset = flag.Bool("reset", false, "Reset the data using the original data files")
var fServe = flag.Bool("serve", false, "Serve the webpage")

func tp(name string) string {
	return "./templates/" + name + ".html"
}

var templates = template.Must(template.ParseFiles(tp("root")))

func rootHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "root.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	flag.Parse()
	if *fReset {
		reset()
	}

	if *fServe {
		db, err := sql.Open("sqlite3", dbPath+dbFilename)
		checkErr(err)
		defer func() {
			err := db.Close()
			checkErr(err)
		}()
		loadNutrients(db)

		http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
			data, err := getIngredientData(db, r.FormValue("ingredients"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			js, err := json.Marshal(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			log.Printf("data is %v and js is %s", data, js)

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		})

		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
		http.HandleFunc("/", rootHandler)
		log.Println("Listening on port :8080")
		http.ListenAndServe(":8080", nil)
	}
}
