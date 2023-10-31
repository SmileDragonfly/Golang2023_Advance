package main

import (
	"fmt"
)

func changeAds(base10 int32) int32 {
	i := 0
	for ; i < 32; i++ {
		val1 := int64(base10) << i
		val2 := int64(1) << 32
		if (val1 & val2) != 0 {
			break
		}
	}
	reverted := ^base10
	var ret int32 = 0
	fmt.Println(i)
	for j := 0; j < 32-i+1; j++ {
		bitVal := reverted & (1 << j)
		ret += bitVal
	}
	return ret
}

func vowelsubstring(s string) int64 {    // Write your code here
	var subs []string
	for i := 0; i < len(s); i++{
		for j:= i+4; j < len(s); j++ {
			sub := s[i:j+1]
			subs=append(subs, sub)
		}
	}
	nRet := 0
	for _, v := range subs {
		var checkMap = map[int32]bool{'a': false,'e':false,'i':false,'o':false,'u':false}
		for _, vs := range v {
			checkMap[vs] = true
		}
	}

	for _, val := range checkMap {
		if !val {
			bCheck = false
		}
	}
	if bCheck {
		nRet++
	}
}
return int64(nRet)
}
func main() {
	ret := vowelsubstring("aaeiouxa")
	fmt.Println(ret)
}
