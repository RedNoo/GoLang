package main

import (
	"fmt"
	"html/template"
	"os"
)

var tpl *template.Template

type person struct {
	Name    string
	Surname string
}

func init() {

	tpl = template.Must(template.ParseFiles("struct.gohtml"))
}

func main() {

	s := person{Name: "john", Surname: "doe"}

	err := tpl.ExecuteTemplate(os.Stdout, "struct.gohtml", s)

	if err != nil {
		fmt.Println(err)
	}

}
