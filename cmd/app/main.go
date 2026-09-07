package main

import (
	"fmt"
	"net/http"

	"AGoTTHT/internal/handler"
)

func main() {
	// 1. Serve the static directory shown in image_4dbb04.png
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 2. Register application routes
	http.HandleFunc("/", handler.HandleHome)
	http.HandleFunc("/grades", handler.HandleGrades)
	http.HandleFunc("/progress", handler.HandleProgress)

	// 3. Start the server
	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
