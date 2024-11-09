package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type data struct {
	Name  string
	Motto string
}

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	var err error
	tpl, err = template.New("").Funcs(fm).ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) < 3 {
		return s // Return as is if the string length is less than 3
	}
	return s[:3] // Return the first 3 characters
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
