package main

import (
	"log"
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
	reloadViewRoot = RELOAD_URL + ":" + config.Server.Port + viewPrefix
}

func main() {
	log.Println("Yellow Boy Begin to run: platform:", runtime.GOOS, ",open port:", config.Server.Port, ",request prifex:", config.Server.Prefix, ",static page in ", config.Assets.Root+STATIC_PAGE_DIR)

	//staticViewHandle := http.StripPrefix("/view/", http.FileServer(http.Dir("./assets/static/")))
	//http.Handle("/view/", staticViewHandle)
	//http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("./template"))))
	//handle := http.FileServer(http.Dir("./template"))
	startLocalServer();
}
