package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func connectDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@127.0.0.1/alura_store?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}

type Product struct {
	Id 					int
	Name  			string
	Description string
	Price 			float64
	Quantity 		int
}

var tpl = template.Must(template.ParseGlob("templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	result, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for result.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = result.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	tpl.ExecuteTemplate(w, "Index", products)
	defer db.Close()
}

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
