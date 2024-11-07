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
	religeon := []string{"Buddha", "Jesus", "Mohamed", "Gandhi", "MLK"}
	err := tpl.Execute(os.Stdout, religeon)
	if err != nil {
		log.Fatalln(err)
	}

}
