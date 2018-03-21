package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/gorilla/mux"
)

//NewRouter creates the router
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/hehe/api/payPage", payPage).Methods("GET")
	r.HandleFunc(apiPrefix+"todos/{posters}", posterPage).Methods("GET")

	r.NotFoundHandler = http.HandlerFunc(NotFound)
	return r
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

//to pay page
func payPage(w http.ResponseWriter, r *http.Request) {
	err := open(reloadViewRoot + config.Assets.Page.Pay)
	if err != nil {
		log.Println("Open url error, the message is ", err)
	}
}

func videoPage(w http.ResponseWriter, r *http.Request) {

	err := open(reloadViewRoot + config.Assets.Page.Video)
	if err != nil {
		log.Println("Open url error, the message is ", err)
	}
}

func posterPage(w http.ResponseWriter, r *http.Request) {
	log.Println("in method poster begin")
	open("")
	log.Println("in method poster end")
}

func open(uri string) error {
	log.Println(config.Assets.Chrome, uri, DISABLE_TRANSLATE, TEST_TYPE, DISABLE_WEB_SECURITY, FULL_SCREEN_PARAM)
	cmd := exec.Command(config.Assets.Chrome, uri, DISABLE_TRANSLATE, TEST_TYPE, DISABLE_WEB_SECURITY, FULL_SCREEN_PARAM)
	return cmd.Start()
}
