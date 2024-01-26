package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"sync"
	tms_proto_licensing "testtms/tms.proto.licensing"
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

func CallCheckLicense(conn *grpc.ClientConn, customerId string, atmId string) error {
	// Call grpc CheckLicense
	req := tms_proto_licensing.CheckLicenseRequest{
		AtmId:      atmId,
		CustomerId: customerId,
	}
	client := tms_proto_licensing.NewLicenseServiceClient(conn)
	resp, err := client.CheckLicense(context.Background(), &req)
	if err != nil {
		log.Println(fmt.Sprintf("[%s][%s]Error: %v", customerId, atmId, err))
		return err
	}
	log.Println(fmt.Sprintf("[%s][%s]Status: %v", customerId, atmId, resp.Status))
	return nil
}

type dbConfig struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

func CallCheckLicenseCustomer(url string, customerId string, dbConf dbConfig) error {
	// Get all atmId from db
	conn, err := gorm.Open(sqlserver.Open(fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		dbConf.user, dbConf.password, dbConf.host, dbConf.port, dbConf.dbname)))
	if err != nil {
		log.Println(err)
		return err
	}
	var atmIds []string
	conn.Table("tblLicenseAtm").Select("AtmId").Where("LicenseId = ?", "73A490B5-5E5F-4496-9162-C38E5C3C21BA").Find(&atmIds)
	log.Println(fmt.Sprintf("Total atmIds: %d", len(atmIds)))
	log.Println(fmt.Sprintf("List atmIds: %+v", atmIds))
	grpcConn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Println(fmt.Sprintf("[%s]Error: %v", customerId, err))
		return err
	}
	// Call grpc
	var wg sync.WaitGroup
	loop := 1
	countAtm := 0
	for i := 0; i < loop; i++ {
		for _, atmId := range atmIds {
			go func(customer string, atm string) {
				wg.Add(1)
				CallCheckLicense(grpcConn, customer, atm)
				wg.Done()
			}(customerId, atmId)
			countAtm++
			if countAtm > 600 {
				break
			}
		}
	}
	wg.Wait()
	return nil
}
