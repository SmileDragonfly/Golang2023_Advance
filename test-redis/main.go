package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
)

func main() {
	// 1. Load token from db
	gormDB, err := gorm.Open(sqlserver.Open("server=192.168.68.242;user id=sa;password=123@123A;port=1433;database=PayGreen_Dev_Hub"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	terminalID := "4200FB47-AE00-4841-B6F5-AC97B9ABAADB"
	token := LoadTokenFromDB(gormDB, terminalID)
	if token == "" {
		log.Fatal("Token not found")
	}
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.68.242:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	err = client.Set(ctx, terminalID, token, 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := client.Get(ctx, terminalID).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("terminalID", val)
}
