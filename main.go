package main

import (
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.htm"))

		data := map[string][]Film{
			"Films": {
				{"The Shawshank Redemption", "Frank Darabont"},
				{"The Godfather", "Francis Ford Coppola"},
				{"The Godfather: Part II", "Francis Ford Coppola"},
				{"The Dark Knight", "Christopher Nolan"}},
		}

		tmpl.Execute(w, data)
	}

	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
