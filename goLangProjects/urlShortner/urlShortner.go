package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func main() {
	jsonPaths := `{
	"/fb": "https://www.facebook.com/",
	"/yt": "https://www.youtube.com/"
}`
	str := gjson.Get(jsonPaths, "/yt")

	fmt.Println(str)
}
