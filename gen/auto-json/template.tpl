package main

//This code has been generated by go generate.
//DO NOT EDIT BY HAND

//{{Title .Name}} resource
type {{Title .Name}} struct{
    {{- range $jsonName, $val := .Fields}}
        {{(Title $jsonName)}} {{TypeOf $val}} `json:"{{$jsonName}}"`{{end}}
}