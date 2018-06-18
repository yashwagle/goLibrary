package main

import(
  "net/http"
  "fmt"
)



func main(){
  h:= http.NewServeMux()

  h.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
  })

  h.HandleFunc("/earth", func(w http.ResponseWriter, r *http.Request) {




  // resp:=string(p)
    //fmt.Fprintf(w, "Hellow earth")
    http.Redirect(w, r, "https://github.com/gophercises", 301)
  })

  h.HandleFunc("/aliens", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello %s",r.RequestURI)
  })

  fmt.Println("Starting server")
  err:=http.ListenAndServe(":9090", h)
  fmt.Println(err.Error())

}
