package main

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"time"
)

type ITest interface {
	DoSomething()
}

func main() {
	test := ITest(nil)
	test.DoSomething()
	// 1. Open DB
	gormDB, err := gorm.Open(sqlserver.Open("server=192.168.68.242;user id=sa;password=123@123A;port=1433;database=PayGreen_Dev_Hub"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// Insert transaction lock
	tx := gormDB.Create(&HubTxnLock{
		Id:            uuid.New().String(),
		TerminalHubId: "9004262D-E86C-4C73-8D5E-28D8FFB9C040",
		HubTID:        "T111111",
		TerminalID:    "",
		CreatedDate:   time.Now(),
		ExpiredTime:   time.Now().Add(time.Minute * 5),
		LockType:      0,
	})
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}
	// Comment: Code for test token
	//terminalID := "4200FB47-AE00-4841-B6F5-AC97B9ABAADB"
	//token := LoadTokenFromDB(gormDB, terminalID)
	//if token == "" {
	//	log.Fatal("Token not found")
	//}
	//client := redis.NewClient(&redis.Options{
	//	Addr:     "192.168.68.242:6379",
	//	Password: "",
	//	DB:       0,
	//})
	//ctx := context.Background()
	//err = client.Set(ctx, terminalID, token, 0).Err()
	//if err != nil {
	//	panic(err)
	//}
	//val, err := client.Get(ctx, terminalID).Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("terminalID", val)
}
