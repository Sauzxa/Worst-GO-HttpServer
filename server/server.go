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
	httpServer := newHTTPLogServer()

	router := mux.NewRouter()
	router.HandleFunc("/", httpServer.handleProducer).Methods("POST")
	router.HandleFunc("/", httpServer.handleConsumer).Methods("GET")

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
