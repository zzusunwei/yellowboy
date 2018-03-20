package main

import (
	"log"
    "fmt"
    "os/exec"
    "runtime"
)

var commands = map[string]string{
    "windows": "cmd /c start",
    "darwin":  "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
    "linux":   "xdg-open",
}
var fullScreenParam = "-kiosk"

// Open calls the OS default program for uri
func Open(uri string) error {
	log.Println("Open ", uri, "in", runtime.GOOS)
    run, ok := commands[runtime.GOOS]
    if !ok {
        return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}
	log.Println(run, uri, fullScreenParam)
    cmd := exec.Command(run, uri, fullScreenParam)
    return cmd.Start()
}