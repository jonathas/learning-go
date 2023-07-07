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

	defer db.Close()

	return products
}
