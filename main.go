package main

import (
	"log"
    "net/http"
)

var (
	version   = "unset"
	build     = "undefined"
	buildDate = "unknown"
    address   = ":8000"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
}

func main() {
	log.SetPrefix("The Tangram Composer ")
    log.Printf("version: %s, build: %s, build date: %s", version, build, buildDate)
    http.HandleFunc("/health", healthCheck)
    log.Printf("Listening on %s", address)
    log.Fatal(http.ListenAndServe(address, nil))
}
