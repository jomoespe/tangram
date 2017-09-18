package router

import (
	"errors"
	"net/http"

	"golang.org/x/net/html"
)

type Configuration struct {
	Timeout int
	Routes  [1]Route // TODO do this as an array
}

type Route struct {
	Service string
	Path    string
}

const (
	dataLocationAttr = "data-loc"
)

var (
	conf Configuration
)

func (conf Configuration) Register() {
	for _, route := range conf.Routes {
		http.HandleFunc(route.Path, route.route)
	}
}

func (route Route) route(w http.ResponseWriter, r *http.Request) {
	node, err := getNode(route.Service)
	if err == nil {
		node := processNode(node)
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		html.Render(w, &node)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func getNode(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &html.Node{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return &html.Node{}, errors.New("url " + url + "returns status " + resp.Status)
	}
	node, err := html.Parse(resp.Body)
	if err != nil {
		return &html.Node{}, err
	}
	return node, nil
}

func processNode(node *html.Node) html.Node {
	if node.Type == html.ElementNode {
		if isHolder, target := isComponentHolder(node); isHolder == true {
			component, err := getNode(target)
			if err == nil {
				cleanNode(node)
				processed := processNode(component)
				node.AppendChild(&processed)
			} else {

			}

		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		processNode(c)
	}
	return *node
}

func isComponentHolder(node *html.Node) (bool, string) {
	for _, a := range node.Attr {
		if a.Key == dataLocationAttr {
			return true, a.Val
		}
	}
	return false, ""
}

func cleanNode(node *html.Node) {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		node.RemoveChild(c)
	}
	// here data-loc attribute should be removed
}
