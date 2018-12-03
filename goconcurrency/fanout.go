package main

import (
	"fmt"
	"sync"
	"time"

	"math/rand"
)

func main() {
	f1 := make(chan int)
	f2 := make(chan int)

	go send(f1)

	go fanOut(f1, f2)

	for v := range f2 {
		fmt.Println(v)
	}

	fmt.Println("exit")
}

func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}

	close(c)
}

func fanOut(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(v2 int) {
			c2 <- timeConsimingWork(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func timeConsimingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)

}
