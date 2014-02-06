package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var isLogging = true

func main() {
	app := httputil.NewSingleHostReverseProxy(&url.URL{Scheme: "http", Host: "localhost:9090", Path: "/"})
	http.HandleFunc("/app/", handler(app))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if isLogging {
			log.Printf("- %v, %v, %v", r.RemoteAddr, r.Method, r.RequestURI)
		}
		p.ServeHTTP(w, r)
	}
}
