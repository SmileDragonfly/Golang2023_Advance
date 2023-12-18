package main

import (
	"fmt"
	"os"
	"os/signal"
)

type Data struct {
	val1 int64
	val2 int64
}

func main() {
	a := new(Data)
	fmt.Printf("a=%p", a)
	fmt.Println("*a=", *a)
	go func() {
		a.val1 = 5
		a.val2 = 10
		fmt.Printf("a=%p", a)
		fmt.Println("*a=", *a)
	}()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
	<-sig
}
