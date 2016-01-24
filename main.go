package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var fLoad = flag.Bool("load", false, "Load the data from the original data files")
var fServe = flag.Bool("serve", false, "Serve the webpage")

func tp(name string) string {
	return "./templates/" + name + ".html"
}

var templates = template.Must(template.ParseFiles(tp("root")))

type Data struct {
	Ingredients string
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{Ingredients: r.FormValue("ingredients")}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "root.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	flag.Parse()
	if *fLoad {
		loadData()
	}

	if *fServe {
		http.HandleFunc("/process", processHandler)
		http.HandleFunc("/", rootHandler)
		http.ListenAndServe(":8080", nil)
		log.Println("Listening on port 8080")
	}
}
