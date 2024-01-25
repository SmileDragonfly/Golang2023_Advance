package main

import "sync"

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go CallTMSAPIGetCustomer("http://192.168.68.252:5050/api/customer", &wg)
	}
	wg.Wait()
}
