package main

import (
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"
)
const (
	STATIC_PAGE_DIR = "./assets/static/"
	URL_PREFIX = "/assets/"
    UPLOAD_DIR   = "./assets/videos/"
	TEMPLATE_DIR = "./assets/posterPages/"
	LOCAL_URL = "http://localhost"
	LOCAL_PORT = ":8080"
	PAY_PAGE = "pay.html"
	VIDEO_PAGE = "video.html"
)

//Todo struct to todo
type Todo struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `json:"name"`
	Completed bool          `json:"completed"`
	Created   time.Time     `json:"createdon"`
}

func main() {
	t := time.Now()
	log.Println("Begin to run: ", t.Format("2006-01-02:15:04:05"))
	http.Handle(URL_PREFIX, http.StripPrefix(URL_PREFIX,http.FileServer(http.Dir(STATIC_PAGE_DIR))))
	router := NewRouter()
	log.Fatal(http.ListenAndServe(LOCAL_PORT, router))
	log.Println("Shutdown at: ", t.Format("2006-01-02:15:04:05"))
}