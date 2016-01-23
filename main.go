package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var fLoad = flag.Bool("load", false, "Load the data from the original data files")

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

	http.HandleFunc("/process", processHandler)
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}
