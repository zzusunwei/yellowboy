package main
import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"os"

	"github.com/gorilla/mux"
)
//NewRouter creates the router
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/payPage", payPage).Methods("GET")
	r.HandleFunc("/api/todos/{posters}", posterPage).Methods("GET")

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
	err := Open(LOCAL_URL+LOCAL_PORT+PAY_PAGE)
	if err != nil {
		log.Println("Open url error, the message is ", err)
	}
}

func videoPage(w http.ResponseWriter, r *http.Request) {
	
	err := Open(LOCAL_URL+LOCAL_PORT+VIDEO_PAGE)
	if err != nil {
		log.Println("Open url error, the message is ", err)
	}
}

func posterPage(w http.ResponseWriter, r *http.Request) {
	log.Println("in method poster begin")
	Open("")
	log.Println("in method poster end")
}
func indexPage(w http.ResponseWriter, r *http.Request){
	log.Println("in method index page")
	pageName := mux.Vars(r)["pageName"];
	var filePath string
	if pageName == "" {
		filePath = STATIC_PAGE_DIR + "index.html"
	}
	filePath = STATIC_PAGE_DIR + mux.Vars(r)["pageName"]
    if exists := isExists(filePath); !exists {
		log.Println("Page not found", filePath)
		http.NotFound(w,r)
		return
	}
	log.Println("To page", filePath)
    http.ServeFile(w,r,filePath)
}

func isExists(path string) bool {
    _, err := os.Stat(path)
    if err == nil {
        return true
    }
    return os.IsExist(err)
}