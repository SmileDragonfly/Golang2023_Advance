package main

import "context"

type UserRepo interface {
	GetAllUsers(ctx context.Context) ([]User, error)
}
