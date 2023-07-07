package controllers

import (
	"net/http"
	"text/template"

	"github.com/jonathas/learning-go/webapp/models"
)

var tpl = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "Index", models.Product{}.GetAll())
}
