package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

//go:generate go run gen.go

func main() {
	var resp map[string]interface{}
	//in, _ := os.Open("user.json")
	b, err := ioutil.ReadFile("user.json")
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(b, &resp); err != nil {
		panic(err)
	}
	data := struct {
		Name   string
		Fields map[string]interface{}
	}{
		Name:   "user",
		Fields: resp,
	}
	tpl, err := template.New("template.tpl").Funcs(template.FuncMap{
		"Title": strings.Title,
		"TypeOf": func(v interface{}) string {
			if v == nil {
				return "string"
			}
			return strings.ToLower(reflect.TypeOf(v).String())
		},
	}).ParseFiles("template.tpl")
	if err != nil {
		panic(err)
	}
	out, _ := os.Create("user.gen.go")
	defer out.Close()
	tpl.Execute(out, data)
}
