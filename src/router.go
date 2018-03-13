package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//NotFound responses to routes not defined
func NotFound(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s\t%s\t%s\t%s\t%d\t%d\t%d",
		r.RemoteAddr,
		r.Method,
		r.RequestURI,
		r.Proto,
		http.StatusNotFound,
		0,
		0,
	)
	w.WriteHeader(http.StatusNotFound)
}

//NewRouter creates the router
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/todos", TodoIndex).Methods("GET")
	r.HandleFunc("/api/todos/{todoID}", TodoShow).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(NotFound)
	return r
}
