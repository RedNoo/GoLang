package main

import (
	"fmt"
	"reflect"
)

func main(){
	fmt.Println(reflect.TypeOf(33));
	fmt.Println(reflect.TypeOf("Hello"))
}