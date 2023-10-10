package main

import (
	"fmt"
	"sync"
	"time"
)

var val int

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	startTime := time.Now()
	for i := 0; i < 1000; i++ {
		go func(counter int) {
			wg.Add(1)
			mutex.Lock()
			val++
			mutex.Unlock()
			fmt.Println("I-Value:", counter, "-", val)
			wg.Done()
		}(i)
	}
	wg.Wait()
	timeElapsed := time.Now().Sub(startTime)
	fmt.Println("Last-Value:", val)
	fmt.Println("Time-Elapsed:", timeElapsed.Milliseconds())
}
