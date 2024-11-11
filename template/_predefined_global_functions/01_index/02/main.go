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
	xs := []string{"zero", "one", "two", "three", "four", "five"}

	data := struct {
		Word  []string
		Lname string
	}{
		xs,
		"Nathnael",
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

}
