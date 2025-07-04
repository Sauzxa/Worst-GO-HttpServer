package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	// Specify the port to listen on
	addr := ":" + port
	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Define routes with specified HTTP methods
	router.HandleFunc("/", handleRoot).Methods("GET") // specify the method
	router.HandleFunc("/hello", handleHello).Methods("GET")

	// Create a new HTTP server
	server := &http.Server{
		Addr:    addr,
		Handler: router, // Use Gorilla Mux router as the handler
	}

	// Start the server
	fmt.Printf("Server listening on %s\n", addr)
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) { // pointer 3la struct to get all data from req url , headers etc ..
	fmt.Fprintf(w, "Welcome to the homepage!") // w instance ml interface to get all methods status code , headers
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
