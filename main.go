package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Set port, default to 8080 if not specified
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Define HTTP handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	// Start the server
	addr := ":" + port
	log.Printf("Starting server on %s...", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
