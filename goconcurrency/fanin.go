package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)

	go receive(even, odd, fanin)

	fmt.Println(fanin)

	for v := range fanin {
		fmt.Println(v)

	}
	fmt.Println("exit")
}

func send(even, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
			fmt.Println("i", i)
		} else {
			odd <- i
			fmt.Println("i", i)
		}

	}
	close(even)
	close(odd)
}

func receive(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fanin <- v
			fmt.Println("even", v)
		}
		wg.Done()
	}()
	go func() {
		for v := range odd {
			fanin <- v
			fmt.Println("odd", v)
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanin)

}
