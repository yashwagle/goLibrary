package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func main() {

	jsonData := `{path:'/yt',url:'https://www.youtube.com'},{path:'/fb',url:'https://www.facebook.com'},{path:'/bpm',url:'http://localhost:8120/amxadministrator/loginForm.jsp'}`

	res := gjson.Get(jsonData, `path`)

	fmt.Println(res)
}
