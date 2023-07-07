package main

import (
	"html/template"
	"net/http"

	"github.com/jonathas/learning-go/webapp/models"

	_ "github.com/lib/pq"
)

var tpl = template.Must(template.ParseGlob("templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "Index", models.Product{}.GetAll())
}

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
