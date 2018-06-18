package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshortgophers": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
    "/mygit" : "https://github.com/yashwagle/goLibrary",

  }
	mapHandler := MapHandlerFunc(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback

	fmt.Println("Starting the server on :9999")
	http.ListenAndServe(":9999", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}


func MapHandlerFunc(pathstoUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := pathstoUrls[r.RequestURI]
		if res == "" {
			fallback.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, res, 301)
		}
	})
}
