package main

import (
	"math/rand"
	"time"
)

func RandomArray(len int, rank int) []int {
	rand.Seed(time.Now().UnixNano())
	ret := make([]int, len)
	for i := range ret {
		ret[i] = rand.Int() % rank
	}
	return ret
}
