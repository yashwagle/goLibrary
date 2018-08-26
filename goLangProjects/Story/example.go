package main
import (
"fmt"
"log"
"net/http"
)
func hello(w http.ResponseWriter, r *http.Request) {
if r.RequestURI == "/form" {
http.ServeFile(w, r, "example.html")
}
switch r.Method {
case "POST":
  if err := r.ParseForm();
   err != nil {
    fmt.Fprintf(w, "ParseForm() err: %v", err)
    return
  }
  fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
  name := r.FormValue("name")
  address := r.FormValue("address")
  fmt.Fprintf(w, "Name = %s\n", name)
  fmt.Fprintf(w, "Address = %s\n", address)
default:
  http.Error(w, "404 not found.", http.StatusNotFound)
}
}
func main() {
  http.HandleFunc("/", hello)
  fmt.Printf("Starting server for testing HTTP POST...\n")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
  }
}
