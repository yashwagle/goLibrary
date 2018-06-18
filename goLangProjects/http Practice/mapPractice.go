package  main

import(
  "fmt"
)


func main(){
  m:=make(map[string]string)
  m["GOT"]="Tyrion"
  m["BB"]="Walter"

  fmt.Println(m["GOT"])
  if m["Narcos"]== ""{
    fmt.Println(m["BB"])
}
}
