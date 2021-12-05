package handlers

import (
	"net/http"

	"github.com/jonathas/learning-go/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.go.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.go.tmpl")
}
