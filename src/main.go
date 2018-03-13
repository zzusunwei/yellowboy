package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter() // this func is in router.go
	log.Fatal(http.ListenAndServe(":8080", router))
}
