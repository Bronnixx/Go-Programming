package main

import (
	"embed"
	"os"
	"text/template"
)
type Person struct{
	Name string
}
var (
	//go:embed templates
	f embed.FS
)
func main(){
	p := Person{"Erick"}
	tmpl, err := template.ParseFS(f, "templates/template.txt")
	if err != nil{
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, p)
	if err != nil{
		panic(err)
	}
}