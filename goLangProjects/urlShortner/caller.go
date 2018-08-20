package main

import (
	"fmt"
	"net/http"

	"github.com/tidwall/gjson"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshortgophers": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":      "https://godoc.org/gopkg.in/yaml.v2",
		"/mygit":           "https://github.com/yashwagle/goLibrary",
	}
	mapHandler := MapHandlerFunc(pathsToUrls, mux)

	jsonPaths := `{
	"/fb": "https://www.facebook.com/",
	"/yt": "https://www.youtube.com/",
	"/go-doc":"https://golang.org/doc/",
	"/jsonlib":"https://github.com/tidwall/gjson"`

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	jsonHandler := JsonHandlerFunc(jsonPaths, mapHandler)

	jsonData := `[{path:'/yt',url:'https://www.youtube.com'},{path:'/fb',url:'https://www.facebook.com'},,{path:'/bpm',url:'http://localhost:8120/amxadministrator/loginForm.jsp'}]`

	fmt.Println("Starting the server on :9999")
	http.ListenAndServe(":9999", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "404 URL not found in the list")
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

func JsonHandlerFunc(pathstoUrls string, fallback http.Handler) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := gjson.Get(pathstoUrls, r.RequestURI)
		if !res.Exists() {
			fallback.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, res.String(), 301)
		}
	})
}
