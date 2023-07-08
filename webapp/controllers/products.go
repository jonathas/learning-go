package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/jonathas/learning-go/webapp/models"
)

var tpl = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "Index", models.Product{}.GetAll())
}

func New(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", 301)
		return
	}

	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")
	quantity := r.FormValue("quantity")

	convertedPrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Println("Error converting price:", err)
	}

	convertedQuantity, err := strconv.Atoi(quantity)
	if err != nil {
		log.Println("Error converting quantity:", err)
	}

	models.Product{}.Insert(name, description, convertedPrice, convertedQuantity)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	tpl.ExecuteTemplate(w, "Edit", models.Product{}.GetById(id))
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", 301)
		return
	}

	id := r.FormValue("id")
	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")
	quantity := r.FormValue("quantity")

	convertedId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error converting id:", err)
	}

	convertedPrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Println("Error converting price:", err)
	}

	convertedQuantity, err := strconv.Atoi(quantity)
	if err != nil {
		log.Println("Error converting quantity:", err)
	}

	models.Product{}.Update(convertedId, name, description, convertedPrice, convertedQuantity)

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	models.Product{}.Delete(id)

	http.Redirect(w, r, "/", 301)
}
