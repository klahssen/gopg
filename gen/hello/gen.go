package main

import (
	"html/template"
	"os"
)

//go:generate go run gen.go

var tpl = template.Must(template.New("").Parse("Hello {{.Name}}"))

func main() {
	data := struct{ Name string }{Name: "John Doe"}
	out, _ := os.Create("output.txt")
	defer out.Close()
	tpl.Execute(out, data)
}
