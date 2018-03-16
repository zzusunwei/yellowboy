package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
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
	serialPort()
	router := NewRouter() // this func is in router.go
	log.Fatal(http.ListenAndServe(":8080", router))
	log.Println("Shutdown at: ", t.Format("2006-01-02:15:04:05"))
}

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

//JSONResponse function to help in responses
func JSONResponse(w http.ResponseWriter, r *http.Request, start time.Time, response []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	log.Printf("%s\t%s\t%s\t%s\t%d\t%d\t%s",
		r.RemoteAddr,
		r.Method,
		r.RequestURI,
		r.Proto,
		code,
		len(response),
		time.Since(start),
	)
	if string(response) != "" {
		w.Write(response)
	}
}

//JSONError function to help in error responses
func JSONError(w http.ResponseWriter, r *http.Request, start time.Time, message string, code int) {
	j := map[string]string{"message": message}
	response, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	log.Printf("%s\t%s\t%s\t%s\t%d\t%d\t%s",
		r.RemoteAddr,
		r.Method,
		r.RequestURI,
		r.Proto,
		code,
		len(response),
		time.Since(start),
	)
	w.Write(response)
}

//TodoIndex handler to route index
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("in method TodoIndex")

	start := time.Now()
	var todos []Todo
	response, err := json.MarshalIndent(todos, "", "    ")
	if err != nil {
		panic(err)
	}
	JSONResponse(w, r, start, response, http.StatusOK)
}

//TodoShow handler to show all todos
func TodoShow(w http.ResponseWriter, r *http.Request) {
	log.Println("in method TodoIndex")
}
