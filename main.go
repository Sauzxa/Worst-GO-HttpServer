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
	addr := ":" + port // using local ipv4 address or localhost address
	log.Printf("Starting server on %s...", addr)
	err := http.ListenAndServe(addr, nil) // create + start server on addr ou tasta9bal reqs tcp fl port ta3 address
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
