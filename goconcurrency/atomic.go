package main

import (
	"fmt"
	"sync"

	"sync/atomic"
)

func main() {

	var wg sync.WaitGroup

	var counter int32 = 0
	gNum := 10000
	wg.Add(gNum)
	for i := 0; i < gNum; i++ {
		go func() {
			atomic.AddInt32(&counter, 1)

			wg.Done()
		}()

	}

	wg.Wait()
	fmt.Println("counter", atomic.LoadInt32(&counter))
}
