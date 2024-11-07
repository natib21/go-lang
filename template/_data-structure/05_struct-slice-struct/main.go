package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type man struct {
	Name  string
	Motto string
}
type car struct {
	Model        string
	Manufacturer string
	Doors        int
}
type items struct {
	Wisdom    []man
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	buddha := man{
		Name:  "Buddha",
		Motto: "the belief of no beliefs ",
	}
	gandhi := man{
		Name:  "India",
		Motto: "Be the change",
	}
	mlk := man{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with haters but with love alone is healed",
	}
	jesus := man{
		Name:  "Jesus",
		Motto: "Love All",
	}
	muhammad := man{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good ",
	}
	t := car{
		Manufacturer: "Toyota",
		Model:        "T324",
		Doors:        4,
	}
	l := car{
		Manufacturer: "Lamborgin",
		Model:        "LA21",
		Doors:        2,
	}
	n := car{
		Manufacturer: "Nexus",
		Model:        "NT24",
		Doors:        4,
	}
	f := car{
		Manufacturer: "Ford",
		Model:        "F1-15",
		Doors:        2,
	}

	religion := []man{buddha, gandhi, mlk, jesus, muhammad}
	cars := []car{t, l, n, f}

	data := items{
		Wisdom:    religion,
		Transport: cars,
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
