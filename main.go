package main

import (
	"net/http"
)

type api struct {
	address string
	url     map[string]string
}

func (a *api) HandlePath(method string, path string) string {
	switch method {
	case "GET":
		switch path {
		case "/":
			return "Home"
		case "/posts":
			return "Get all posts"
		default:
			return "404"
		}
	case "POST":
		switch path {
		case "/posts":
			return "Post a post"
		default:
			return "404"
		}
	default:
		return "404"
	}
}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	responseString := a.HandlePath(r.Method, r.URL.Path)
	w.Write([]byte(responseString))
}

func main() {
	s := &api{address: ":8080"}
	http.ListenAndServe(s.address, s)
}
