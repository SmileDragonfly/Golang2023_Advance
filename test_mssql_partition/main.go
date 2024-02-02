package main

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"sync"
)

func main() {
	db, err := gorm.Open(sqlserver.Open("sqlserver://tms:123@123A@192.168.68.31:1433?database=CHUNG_VCB_CUSTOMER"))
	if err != nil {
		panic(err)
	}
	//db.AutoMigrate(&User{})
	//db.AutoMigrate(&DoCommandHistory{})
	// Close the underlying *sql.DB connection when done
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()
	// Insert data
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		go func(wg *sync.WaitGroup) {
			wg.Add(1)
			for j := 0; j < 100000; j++ {
				u := DoCommandHistory{}
				u.RandomData()
				db.Create(&u)
				fmt.Println(u.ID)
			}
			wg.Done()
		}(&wg)
	}
	wg.Wait()
}
