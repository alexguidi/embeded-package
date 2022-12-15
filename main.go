package main

import (
	"embed"
	_ "embed"
)

//go:embed static
var f embed.FS

func main() {

	data, _ := f.ReadFile("static/index.html")
	print(string(data))

}
