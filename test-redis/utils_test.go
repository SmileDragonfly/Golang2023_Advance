package main

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"testing"
)

func BenchmarkLoadTokenFromCache(b *testing.B) {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.68.242:6379",
		Password: "",
		DB:       0,
	})
	for i := 0; i < b.N; i++ {
		LoadTokenFromCache(client, "4200FB47-AE00-4841-B6F5-AC97B9ABAADB")
	}
}

func BenchmarkLoadTokenFromDB(b *testing.B) {
	gormDB, err := gorm.Open(sqlserver.Open("server=192.168.68.242;user id=sa;password=123@123A;port=1433;database=PayGreen_Dev_Hub"), &gorm.Config{})
	if err != nil {
		b.Fatal(err)
	}
	db, err := gormDB.DB()
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()
	for i := 0; i < b.N; i++ {
		LoadTokenFromDB(gormDB, "4200FB47-AE00-4841-B6F5-AC97B9ABAADB")
	}
}

func BenchmarkLoadLockFromCache(b *testing.B) {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.68.242:6379",
		Password: "",
		DB:       0,
	})
	for i := 0; i < b.N; i++ {
		LoadLockFromCache(client, "4200FB47-AE00-4841-B6F5-AC97B9ABAADB")
	}
}

func BenchmarkLoadLockFromDB(b *testing.B) {
	gormDB, err := gorm.Open(sqlserver.Open("server=192.168.68.242;user id=sa;password=123@123A;port=1433;database=PayGreen_Dev_Hub"), &gorm.Config{})
	if err != nil {
		b.Fatal(err)
	}
	db, err := gormDB.DB()
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()
	for i := 0; i < b.N; i++ {
		LoadLockFromDB(gormDB, "4200FB47-AE00-4841-B6F5-AC97B9ABAADB")
	}
}
