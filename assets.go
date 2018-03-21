package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port   string
		Prefix string
	}
	Assets struct {
		Chrome string
		Root   string
		Page   struct {
			Pay   string
			Video string
		}
	}
}

func loadConfig(file string) Config {
	log.Println("try to load config file in", file)
	yellowBoyConf := Config{}
	if isExt := isExists(file); !isExt {
		log.Fatal("No config file found")
	}
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}
	err = yaml.Unmarshal(b, &yellowBoyConf)
	if err != nil {
		log.Fatal()
	}
	log.Println("load Config:", yellowBoyConf)
	return yellowBoyConf
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}
