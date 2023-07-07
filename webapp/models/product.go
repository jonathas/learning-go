package models

import "github.com/jonathas/learning-go/webapp/db"

type Product struct {
	Id 					int
	Name  			string
	Description string
	Price 			float64
	Quantity 		int
}

func (product Product) GetAll() []Product {
	db := db.ConnectDB()
	result, err := db.Query("SELECT * FROM products")
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

	defer db.Close()

	return products
}

func (product Product) Insert(name, description string, price float64, quantity int) {
	db := db.ConnectDB()

	insert, err := db.Prepare("INSERT INTO products(name, description, price, quantity) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(name, description, price, quantity)

	defer db.Close()
}

func (product Product) GetById(id string) Product {
	db := db.ConnectDB()

	result, err := db.Query("SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	p := Product{}

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
	}

	defer db.Close()

	return p
}

func (product Product) Delete(id string) {
	db := db.ConnectDB()

	delete, err := db.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)

	defer db.Close()
}
