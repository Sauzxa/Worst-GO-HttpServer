package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type httpLogServer struct {
	Log *Log
}

func newHTTPLogServer() *httpLogServer {
	return &httpLogServer{Log: NewLog()}
}

func NewHTTPServer(addr string) *http.Server {
	router := mux.NewRouter()
	router.HandleFunc("/", server.handleProducer).Methods("POST")
	router.HandleFunc("/", server.handleConsumer).Methods("GET")

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
