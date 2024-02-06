package main

import (
	"context"
	"gorm.io/gorm"
)

type QueryGorm struct {
	db *gorm.DB
}

func (q *QueryGorm) GetAllUsers(ctx context.Context) ([]User, error) {
	var users []User
	tx := q.db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}
