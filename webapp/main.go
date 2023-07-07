package main

import (
	"net/http"

	"github.com/jonathas/learning-go/webapp/routes"
)

func main() {
	routes.LoadRoutes()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
