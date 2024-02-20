package main

import "time"

func main() {
	ch := make(chan int, 1)
	go func(chan int) {
		time.Sleep(2 * time.Second)
		close(ch)
		time.Sleep(2 * time.Second)
		<-ch
		rev1, ok1 := <-ch
		println(rev1)
		println(ok1)
	}(ch)
	ch <- 10
	ch <- 20
	time.Sleep(5 * time.Second)
}
