package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type data struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	buddha := data{
		Name:  "Buddha",
		Motto: "the belief of no beliefs ",
	}
	gandhi := data{
		Name:  "India",
		Motto: "Be the change",
	}
	mlk := data{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with haters but with love alone is healed",
	}
	jesus := data{
		Name:  "Jesus",
		Motto: "Love All",
	}
	muhammad := data{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good ",
	}

	sage := []data{buddha, gandhi, mlk, jesus, muhammad}

	err := tpl.Execute(os.Stdout, sage)
	if err != nil {
		log.Fatalln(err)
	}
}
