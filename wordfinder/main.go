package main

import (
	"fmt"
	"io/ioutil"
)

//https://yourbasic.org/golang/compare-slices/
func Equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {

	keyword := "Cura"

	content, _ := ioutil.ReadFile("temp1.txt")

	arrKey := []byte(keyword)

	for key, value := range content {

		if value == arrKey[0] {
			arrCheck := content[key : key+len(keyword)]

			if Equal(arrKey, arrCheck) {
				fmt.Printf("dosya içerisinde %d  nolu karakterden itibaren bulundu \n", key)
			}
		}

	}
}
