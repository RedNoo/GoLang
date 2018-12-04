package main

import (
	"fmt"
	"html/template"
	"os"
)

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseFiles("slice.gohtml"))
}

func main() {

	s := []string{"a", "b", "c"}

	err := tpl.ExecuteTemplate(os.Stdout, "slice.gohtml", s)

	if err != nil {
		fmt.Println(err)
	}

}
