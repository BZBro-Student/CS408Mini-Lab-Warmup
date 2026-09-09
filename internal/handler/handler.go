package handler

import (
	"AGoTTHT/internal/api/canvas"
	"AGoTTHT/view/components"
	"AGoTTHT/view/layout"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
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
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".env failed to load")
	}
	token := os.Getenv("CANVAS_API_TOKEN")
	domain := os.Getenv("CANVAS_DOMAIN")

	if domain == "" || token == "" {
		log.Fatal("No value assigned to the environment variables")
	}

	client := canvas.NewClient(domain, token)
	courses, err := client.FetchFavoriteCoursesGrades(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderHelper(w, r, components.Grades(courses))
}

// HandleProgress serves the progress path
func HandleProgress(w http.ResponseWriter, r *http.Request) {
	renderHelper(w, r, components.CourseProgress())
}
