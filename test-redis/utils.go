package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func LoadTokenFromDB(db *gorm.DB, terminalID string) string {
	var token string
	db.Model(&APIToken{}).Where("TerminalID = ?", terminalID).Select("Token").Scan(&token)
	return token
}

func LoadTokenFromCache(cache *redis.Client, terminalID string) string {
	token, err := cache.Get(context.Background(), terminalID).Result()
	if err != nil {
		return ""
	}
	return token
}

func LoadLockFromDB(db *gorm.DB, terminalHubID string) string {
	var lock string
	db.Model(&HubTxnLock{}).Where("TerminalHubId = ?", terminalHubID).Select("Id").Scan(&lock)
	return lock
}

func LoadLockFromCache(cache *redis.Client, terminalHubID string) string {
	lock, err := cache.Get(context.Background(), terminalHubID).Result()
	if err != nil {
		return ""
	}
	return lock
}
