package main

import (
	"log"
	"net/http"
)

type Configuration struct {
	Timeout int
	Routes  Route // TODO do this as an array
}
type Route struct {
	Service string
	Path    string
}

var (
	version   = "unset"
	build     = "undefined"
	buildDate = "unknown"
	address   = ":8000"
	conf      Configuration
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func loadConfiguration() {
	routes := Route{Service: "dachop-index", Path: "/dachop/"}
	conf = Configuration{
		Timeout: 5000,
		Routes:  routes}

}

func main() {
	log.SetPrefix("The Tangram Composer ")
	log.Printf("version: %s, build: %s, build date: %s", version, build, buildDate)
	loadConfiguration()
	log.Printf("configuration loaded: %s", conf)

	http.HandleFunc("/health", healthCheck)
	log.Printf("Listening on %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
