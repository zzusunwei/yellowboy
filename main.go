package main

import (
	"log"
	"net/http"
)

const (
	CONFIG_FILE_PATH  = "./assets/config.yml"
	STATIC_PAGE_DIR   = "static/"
	RELOAD_URL        = "http://localhost"
	FULL_SCREEN_PARAM = "-kiosk"
)

var (
	config         Config
	reloadViewRoot string
)

func init() {
	config = loadConfig(CONFIG_FILE_PATH)
	reloadViewRoot = RELOAD_URL + ":" + config.Server.Port + "/" + config.Server.Prefix + "/"
}

func main() {
	log.Println("Yellow Boy Begin to run: open port:", config.Server.Port, ",request prifex:", config.Server.Prefix, ",static page in ", config.Assets.Root+STATIC_PAGE_DIR)
	http.Handle(config.Server.Prefix, http.StripPrefix(config.Server.Prefix, http.FileServer(http.Dir(config.Assets.Root+STATIC_PAGE_DIR))))
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":"+config.Server.Port, router))
}
