package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

var count int = 0
var countLock sync.Mutex

func CallTMSAPIGetCustomer(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	method := "GET"
	client := &http.Client{}
	client.Timeout = time.Minute
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = body
	countLock.Lock()
	defer countLock.Unlock()
	count++
	fmt.Println(count)
	fmt.Println(string(body))
}
