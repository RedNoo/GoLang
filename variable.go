package main

import "fmt"
import "runtime"

type x int

func main() {

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}
