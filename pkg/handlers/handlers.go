package handlers

import (
	"net/http"
	"github.com/ostap-sulyk/udemygo/pkg/render"
)

// Home handles the home page request.
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About handles the about page request.
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
