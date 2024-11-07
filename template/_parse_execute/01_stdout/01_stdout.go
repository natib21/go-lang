package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	f, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = f.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
