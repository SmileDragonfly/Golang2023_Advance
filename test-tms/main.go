package main

import (
	"sync"
)

func main() {
	//var wg sync.WaitGroup
	//for i := 0; i < 1000; i++ {
	//	wg.Add(1)
	//	go CallTMSAPIGetCustomer("http://192.168.68.252:5050/api/customer", &wg)
	//}
	//wg.Wait()
	wg := sync.WaitGroup{}
	// Get all customer and db information
	// Call http request to tms-api to get list customer
	url := "http://192.168.68.252:5050/api/customer"
	customerList, err := GetCustomerList(url)
	if err != nil {
		panic(err)
	}
	for _, customer := range customerList {
		password, err := DecryptAes(customer.DbPassword)
		if err != nil {
			panic(err)
		}
		go CallCheckLicenseCustomer("127.0.0.1:5003",
			customer.Id,
			dbConfig{
				host:     customer.DbServer,
				port:     customer.DbPort,
				user:     customer.DbUser,
				password: password,
				dbname:   customer.DbName,
			},
			&wg)
	}
	wg.Wait()
}
