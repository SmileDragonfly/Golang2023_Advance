package main

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"sync"
)

func main() {
	db, err := gorm.Open(sqlserver.Open("sqlserver://sa:123@123A@192.168.68.242:1433?database=gormdb"))
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
	for i := 0; i < 100; i++ {
		var wg sync.WaitGroup
		for j := 0; j < 20; j++ {
			go func(wg *sync.WaitGroup) {
				wg.Add(1)
				var us []User
				for k := 0; k < 5000; k++ {
					u := User{}
					u.RandomData()
					us = append(us, u)
				}
				db.CreateInBatches(&us, 100)
				wg.Done()
			}(&wg)
		}
		wg.Wait()
		log.Println("Done loop", i)
	}
}
