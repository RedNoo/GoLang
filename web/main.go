package main

import (
	"fmt"
	"html/template"
	"os"
)

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseFiles("a.gohtml", "b.gohtml"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "a.gohtml", "this is a paragraph -1 ")

	if err != nil {
		fmt.Println(err)
	}

}
