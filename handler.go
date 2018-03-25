package main

import (
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

//NewRouter creates the router
func startLocalServer() {
	staticViewHandle := http.StripPrefix("/hehe/view/", http.FileServer(http.Dir("./assets/static/")))
	http.Handle("/hehe/view/", staticViewHandle)
	http.HandleFunc("/hehe/api/pay", payPage)
	http.HandleFunc("/hehe/api/videoPage", videoPage)
	log.Fatal(http.ListenAndServe(":"+config.Server.Port, nil))
}

//to pay page
func payPage(w http.ResponseWriter, r *http.Request) {
	log.Println("in pay page")
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
	var chromePath string
	switch runtime.GOOS {
		case "windows": 
			chromePath = config.Assets.Chrome.Windows
		case "darwin":  
			chromePath = config.Assets.Chrome.Mac
		case "linux":   
			chromePath = config.Assets.Chrome.Linux
		default:
	}
	log.Println(chromePath, uri, DISABLE_TRANSLATE, TEST_TYPE, DISABLE_WEB_SECURITY, FULL_SCREEN_PARAM)
	cmd := exec.Command(chromePath, uri, DISABLE_TRANSLATE, TEST_TYPE, DISABLE_WEB_SECURITY, FULL_SCREEN_PARAM)
	return cmd.Start()
}
