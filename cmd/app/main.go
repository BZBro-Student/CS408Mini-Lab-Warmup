package main

import (
	"log"
	"net/http"
)

func startServer(port string) {
	log.Printf("Server starting on http://localhost%s\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func main() {
	port := ":8080"
	startServer(port)
}
