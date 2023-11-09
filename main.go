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
	data := map[string][]Film{
		"Films": {
			{"The Shawshank Redemption", "Frank Darabont"},
			{"The Godfather", "Francis Ford Coppola"},
			{"The Godfather: Part II", "Francis Ford Coppola"},
			{"The Dark Knight", "Christopher Nolan"}},
	}

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.htm"))
		tmpl.Execute(w, data)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		newFilm := Film{title, director}
		data["Films"] = append(data["Films"], newFilm)

		newLi := `
		<li class="list-group-item bg-primary text-white">
			{{ .Title }} - {{ .Director }}
		</li>
		`
		tmpl, _ := template.New("t").Parse(newLi)
		tmpl.Execute(w, newFilm)
	}

	http.HandleFunc("/add-film/", h2)
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
