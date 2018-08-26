package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend.html")

	var t = []byte("hello")
	w.Write(t)

}
func start(w http.ResponseWriter, r *http.Request) {

	var t = []byte("hello")

	w.Write(t)

}

func main() {
	http.HandleFunc("/", hello)
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("scripts"))))
	http.HandleFunc("/start", start)
	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal(err)
	}
}
