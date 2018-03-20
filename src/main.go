package main

import (
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Todo struct to todo
type Todo struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `json:"name"`
	Completed bool          `json:"completed"`
	Created   time.Time     `json:"createdon"`
}
var url = "http://localhost"
var port = ":8080"

func main() {
	t := time.Now()
	log.Println("Begin to run: ", t.Format("2006-01-02:15:04:05"))
	router := NewRouter() // this func is in router.go
	log.Fatal(http.ListenAndServe(port, router))
	log.Println("Shutdown at: ", t.Format("2006-01-02:15:04:05"))
}