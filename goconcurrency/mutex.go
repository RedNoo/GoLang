package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var m sync.Mutex
	var counter int = 0
	gNum := 10000
	wg.Add(gNum)
	for i := 0; i < gNum; i++ {
		go func() {
			m.Lock()
			a := counter
			runtime.Gosched()
			a++
			counter = a
			m.Unlock()
			wg.Done()
		}()

	}

	wg.Wait()
	fmt.Println("counter", counter)
}
