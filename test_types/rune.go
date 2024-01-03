package main

import "fmt"

func CreateRune() {
	sRune := "ab£"
	rRune := []rune(sRune)
	bRune := []byte(sRune)
	fmt.Printf("sRune: %s\n", sRune)
	fmt.Printf("sRune UTF-8: %U\n", rRune)
	fmt.Printf("sRune byte: %X\n", bRune)
	fmt.Printf("sRune len: %d\n", len(sRune))
}
