package main

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
func main() {

}
