package main

import (
	"fmt"
	"unsafe"
)

func CreateRune() {
	sRune := "ab£"
	rRune := []rune(sRune)
	bRune := []byte(sRune)
	fmt.Printf("sRune: %s\n", sRune)
	fmt.Printf("sRune UTF-8: %U\n", rRune)
	fmt.Printf("sRune byte: %X\n", bRune)
	fmt.Printf("sRune len: %d\n", len(sRune))
}

func CreateArray() {
	sArray := [4]string{"a", "b"}
	fmt.Printf("sArray: %s\n", sArray)
	fmt.Printf("sArray len: %d\n", len(sArray))
}

func CreateStruct() {
	type TestStruct struct {
		s string
		n int
	}
	sStruct := TestStruct{}
	sStruct.s = "This is a test string in struct"
	sStruct.n = 100
	fmt.Printf("sStruct: %+v\n", sStruct)
	fmt.Printf("Empty struct: %+v\n", unsafe.Sizeof(TestStruct{}))
	fmt.Printf("Struct size: %d\n", unsafe.Sizeof(sStruct))
}
