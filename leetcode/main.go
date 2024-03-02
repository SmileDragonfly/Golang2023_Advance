package main

import (
	"fmt"
)

func main() {
	nums := []int{21, 35, 90, 51, 27, 19, 57}
	k := 90
	oper := findOperation2(nums, k)
	fmt.Print(oper)
}
