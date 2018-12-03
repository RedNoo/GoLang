package main

import "fmt"

func foo(ch1, ch2, q chan int) {
	for i := 0; i < 9; i++ {
		//fmt.Println(i)
		if i%2 == 0 {
			ch1 <- i
		} else {
			ch2 <- i
		}

	}

	close(ch1)
	close(ch2)
	q <- 0
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	quit := make(chan int)

	go foo(c1, c2, quit)

	for {
		select {
		case s := <-c1:
			fmt.Println("c1 kanalından ", s)
		case s := <-c2:
			fmt.Println("c2 kanalından ", s)
		case s := <-quit:

			fmt.Println("çıkış ", s)
			close(quit)
			return
		}
	}
}
