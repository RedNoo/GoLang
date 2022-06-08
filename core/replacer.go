package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main(){
	var str string = "this year ###"
	var year string = strconv.Itoa(time.Now().Year())

	replacer := strings.NewReplacer("###", year)

	result := replacer.Replace(str)

	fmt.Println(result)


}