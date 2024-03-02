package main

import (
	"fmt"
	"sort"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var maxLength int
	var sSub string = string(s[0])
	maxLength = len(sSub)
	for i := 1; i < len(s); i++ {
		for j := 0; j < len(sSub); j++ {
			if s[i] == sSub[j] {
				if maxLength < len(sSub) {
					maxLength = len(sSub)
				}
				sSub = sSub[j+1:]
				break
			}
		}
		sSub += string(s[i])
	}
	return maxLength
}

func findOperation(nums []int, k int) int {
	sort.Ints(nums)
	operations := 0
	for _, v := range nums {
		if v < k {
			operations++
		} else {
			break
		}
	}
	return operations
}

func findOperation2(nums []int, k int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	operations := 0
	for {
		if len(nums) < 2 {
			return operations
		}
		if nums[0] < k || nums[1] < k {
			operations++
			add := nums[0]*2 + nums[1]
			nums = nums[2:]
			index := sort.Search(len(nums), func(i int) bool { return nums[i] >= add })
			nums = append(nums, 0)
			copy(nums[index+1:], nums[index:])
			nums[index] = add
		} else {
			break
		}
	}
	return operations
}

func main() {
	nums := []int{21, 35, 90, 51, 27, 19, 57}
	k := 90
	oper := findOperation2(nums, k)
	fmt.Print(oper)
}
