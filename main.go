package main

import (
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

	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}
