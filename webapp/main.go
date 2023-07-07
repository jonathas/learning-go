package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Name  			string
	Description string
	Price 			float64
	Quantity 		int
}

var tpl = template.Must(template.ParseGlob("templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{"Milk", "Milk is a white liquid produced by the mammary glands of mammals.", 2.99, 10},
		{"Cheese", "Cheese is a dairy product derived from milk that is produced in a wide range of flavors, textures, and forms by coagulation of the milk protein casein.", 5.99, 20},
		{"Yogurt", "Yogurt, yoghurt or yoghourt is a food produced by bacterial fermentation of milk.", 3.99, 5},
		{"Butter", "Butter is a dairy product with high butterfat content which is solid when chilled and at room temperature in some regions, and liquid when warmed.", 4.99, 15},
		{"Cream", "Cream is a dairy product composed of the higher-butterfat layer skimmed from the top of milk before homogenization.", 4.99, 15},
		{"Ice Cream", "Ice cream is a sweetened frozen food typically eaten as a snack or dessert.", 6.99, 25},
		{"Cottage Cheese", "Cottage cheese is a fresh cheese curd product with a mild flavor.", 3.99, 5},
		{"Sour Cream", "Sour cream is a dairy product obtained by fermenting regular cream with certain kinds of lactic acid bacteria.", 3.99, 5},
		{"Whipped Cream", "Whipped cream is cream that is whipped by a whisk or mixer until it is light and fluffy.", 3.99, 5},
		{"Condensed Milk", "Condensed milk is cow's milk from which water has been removed.", 3.99, 5},
		{"Evaporated Milk", "Evaporated milk, known in some countries as unsweetened condensed milk, is a shelf-stable canned milk product with about 60% of the water removed from fresh milk.", 3.99, 5},
		{"Powdered Milk", "Powdered milk or dried milk is a manufactured dairy product made by evaporating milk to dryness.", 3.99, 5},
		{"Chocolate Milk", "Chocolate milk is sweetened chocolate-flavored milk.", 3.99, 5},
	}

	tpl.ExecuteTemplate(w, "Index", products)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
