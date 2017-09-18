package router

import (
    //"bytes"
    "net/http"
    //"fmt"

    "golang.org/x/net/html"
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
    resp, err := http.Get(route.Service)
    if err == nil {
        root, _ := html.Parse(resp.Body)
        node := processNode(root)
        w.Header().Set("Content-Type", "text/html")
        w.WriteHeader(http.StatusOK)
        html.Render(w, &node)
    } else {
        w.WriteHeader(http.StatusNotFound)
    }
}

func (conf Configuration) Register() {
    for _, route := range conf.Routes {
        http.HandleFunc(route.Path, route.route)
    }
}

func processNode(node *html.Node) (html.Node) {
    return *node
}