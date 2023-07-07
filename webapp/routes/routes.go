package routes

import (
	"net/http"

	"github.com/jonathas/learning-go/webapp/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
