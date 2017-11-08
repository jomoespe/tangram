package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jomoespe/tangram/router"
	"github.com/pelletier/go-toml"
)

const (
	configFile = "tangram.toml"
)

var (
	version   = "unset"
	build     = "undefined"
	buildDate = "unknown"
	address   = ":8000"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "Tangram/"+version)
	w.WriteHeader(http.StatusOK)
}

func loadConfig() router.Config {
	doc, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	config := router.Config{}
	toml.Unmarshal(doc, &config)
	return config
}

func main() {
	log.SetPrefix("Tangram ")
	log.Printf("version: %s, build: %s, build date: %s", version, build, buildDate)
	conf := loadConfig()
	conf.Register()
	http.HandleFunc("/health", healthCheck)
	log.Printf("Listening on %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
