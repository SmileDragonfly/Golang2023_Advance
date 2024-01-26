package main

func main() {
	//var wg sync.WaitGroup
	//for i := 0; i < 1000; i++ {
	//	wg.Add(1)
	//	go CallTMSAPIGetCustomer("http://192.168.68.252:5050/api/customer", &wg)
	//}
	//wg.Wait()
	CallCheckLicenseCustomer("127.0.0.1:5003", "6c22bd72-6a9f-420d-8cc2-62282a3eeb5e", dbConfig{
		host:     "192.168.68.31",
		port:     1433,
		user:     "tms",
		password: "123@123A",
		dbname:   "CHUNG_VCB_Customer",
	})
}
