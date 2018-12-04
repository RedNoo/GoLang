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

type color struct {
	Name string
}

func init() {

	tpl = template.Must(template.ParseFiles("nestedtemplate.gohtml", "partial.gohtml"))
}

func main() {

	s1 := person{Name: "john", Surname: "doe"}
	s2 := person{Name: "jane", Surname: "doe"}

	c1 := color{Name: "red"}
	c2 := color{Name: "blue"}

	s := []person{s1, s2}

	c := []color{c1, c2}

	data := struct {
		Colors  []color
		Persons []person
	}{
		c, s,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "nestedtemplate.gohtml", data)

	if err != nil {
		fmt.Println(err)
	}

}
