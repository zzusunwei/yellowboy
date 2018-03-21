package main

import (
	"log"
	"net/http"
	"runtime"
)

const (
	CONFIG_FILE_PATH     = "./assets/config.yml"
	STATIC_PAGE_DIR      = "static/"
	RELOAD_URL           = "http://localhost"
	FULL_SCREEN_PARAM    = "--kiosk"
	DISABLE_TRANSLATE    = "--disable-translate"
	DISABLE_WEB_SECURITY = "--disable-web-security"
	TEST_TYPE            = "--test-type"
)

var (
	config         Config
	reloadViewRoot string
	apiPrefix      string
	viewPrefix     string
)

func init() {
	config = loadConfig(CONFIG_FILE_PATH)
	apiPrefix = config.Server.Prefix + "/api/"
	viewPrefix = config.Server.Prefix + "/view/"
	reloadViewRoot = RELOAD_URL + ":" + config.Server.Port + "/" + viewPrefix
}

func main() {
	log.Println("Yellow Boy Begin to run: platform:", runtime.GOOS, "open port:", config.Server.Port, ",request prifex:", viewPrefix, ",static page in ", config.Assets.Root+STATIC_PAGE_DIR)
	http.Handle(viewPrefix, http.StripPrefix(viewPrefix, http.FileServer(http.Dir(config.Assets.Root+STATIC_PAGE_DIR))))
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":"+config.Server.Port, router))
}
