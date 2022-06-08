package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i <= 10; i++ {
		seconds := time.Now().UnixNano()
		rand.Seed(seconds)
		rint := rand.Intn((100))
		fmt.Println(rint)
	}

}
