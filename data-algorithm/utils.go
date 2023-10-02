package main

import (
	"math/rand"
	"time"
)

func RandomArray(len int) []int {
	rand.Seed(time.Now().UnixNano())
	ret := make([]int, len)
	for i := range ret {
		ret[i] = rand.Int() % 1000
	}
	return ret
}
