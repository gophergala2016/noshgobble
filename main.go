package main

import (
	"flag"

	_ "github.com/mattn/go-sqlite3"
)

var fLoad = flag.Bool("load", false, "Load the data from the original data files")

func main() {
	flag.Parse()
	if *fLoad {
		loadData()
	}
}
