package main

import (
	"fmt"
	"html/template"
	"os"

	"strings"
)

var tpl *template.Template

type person struct {
	Name    string
	Surname string
}

type color struct {
	Name string
}

var fm = template.FuncMap{
	"tl": strings.ToLower,
	"tu": strings.ToUpper,
}

func init() {

	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("functionmapping.gohtml"))
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

	err := tpl.ExecuteTemplate(os.Stdout, "functionmapping.gohtml", data)

	if err != nil {
		fmt.Println(err)
	}

}
