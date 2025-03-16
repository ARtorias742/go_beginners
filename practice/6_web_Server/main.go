package main

import (
	"fmt"
	"log"
	"net/http"

	handlers "github.com/ARtorias742/6_web_server/handlers"
)

func main() {
	// Serve static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Define routes
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/api/users", handlers.UsersHandler)

	// Start the server
	port := ":8080"
	fmt.Printf("Server starting on http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
