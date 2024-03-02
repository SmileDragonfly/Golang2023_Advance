package main

import (
	"testing"
)

func BenchmarkFindOperation2(b *testing.B) {
	nums := []int{21, 35, 90, 51, 27, 19, 57}
	k := 90
	for i := 0; i < b.N; i++ {
		findOperation2(nums, k)
	}
}
