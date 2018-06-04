package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main() {

	// Using slice here to print out in the html file
	entities := []string{"Things", "ThingTemplates", "ThingShapes", "DataShapes"}

	err := tpl.Execute(os.Stdout, entities)
	if err != nil {
		log.Fatal(err)
	}
}
