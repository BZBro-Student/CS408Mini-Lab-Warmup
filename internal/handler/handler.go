package handler

import (
	"AGoTTHT/view/components"
	"AGoTTHT/view/layout"
	"net/http"

	"github.com/a-h/templ"
)

// renderHelper decides whether to send a partial or full page
func renderHelper(w http.ResponseWriter, r *http.Request, component templ.Component) {
	if r.Header.Get("HX-Request") == "true" {
		component.Render(r.Context(), w)
	} else {
		layout.Base(component).Render(r.Context(), w)
	}
}

// HandleHome serves the root path
func HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	renderHelper(w, r, components.Home())
}

// HandleGrades serves the grades path
func HandleGrades(w http.ResponseWriter, r *http.Request) {
	renderHelper(w, r, components.Grades())
}

// HandleProgress serves the progress path
func HandleProgress(w http.ResponseWriter, r *http.Request) {
	renderHelper(w, r, components.Grades())
}
