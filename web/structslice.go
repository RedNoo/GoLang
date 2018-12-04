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

	tpl = template.Must(template.ParseFiles("structclice.gohtml"))
}

func main() {

	s1 := person{Name: "john", Surname: "doe"}
	s2 := person{Name: "jane", Surname: "doe"}

	s := []person{s1, s2}

	err := tpl.ExecuteTemplate(os.Stdout, "structclice.gohtml", s)

	if err != nil {
		fmt.Println(err)
	}

}
