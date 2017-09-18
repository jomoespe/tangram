package main

import (
	"log"
	"net/http"

	"github.com/jomoespe/tangram/router"
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

func createConfiguration() router.Configuration {
	return router.Configuration{
		Timeout: 5000,
		Routes: [...]router.Route{
			router.Route{
				Path:    "/dachop/",
				Service: "http://localhost:81/",
			},
//			router.Route{
//				Path:    "/zooplus/",
//				Service: "http://www.zooplus.es",
//			},
		},
	}
}

func main() {
	log.SetPrefix("The Tangram Composer ")
	log.Printf("version: %s, build: %s, build date: %s", version, build, buildDate)
	conf := createConfiguration()
	conf.Register()
	http.HandleFunc("/health", healthCheck)
	log.Printf("Listening on %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
