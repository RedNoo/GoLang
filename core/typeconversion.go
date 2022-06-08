package main

import (
	"fmt"
	"reflect"
)

func main() {
	var height float64
	height = 2.32
	fmt.Println(reflect.TypeOf(height))
	fmt.Println(reflect.TypeOf(int(height)))
}
