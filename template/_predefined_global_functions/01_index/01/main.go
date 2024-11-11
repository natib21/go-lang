package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	xss := []string{"zero", "one", "two", "three", "four", "five"}
	err := tpl.Execute(os.Stdout, xss)
	if err != nil {
		log.Fatalln(err)
	}

}
