package router

import (
    "bytes"
    "net/http"
    "fmt"
)

var (
    conf Configuration
)

type Configuration struct {
    Timeout int
    Routes  [2]Route // TODO do this as an array
}

type Route struct {
    Service string
    Path    string
}

func (route Route) route(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Here we'll routing %s to %s\n\n", route.Path, route.Service)
    resp, _ := http.Get(route.Service)
    body := new(bytes.Buffer)
    body.ReadFrom(resp.Body)


    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprintln(w, body.String())
}

func (conf Configuration) Register() {
    for _, route := range conf.Routes {
        http.HandleFunc(route.Path, route.route)
    }
}
