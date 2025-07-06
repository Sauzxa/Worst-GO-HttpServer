package main

import (
	"log"

	"github.com/yourusername/worst-go-httpserver/server"
)

func main() {
	svr := server.NewHTTPServer(":8080")
	log.Fatal(svr.ListenAndServe())
}
