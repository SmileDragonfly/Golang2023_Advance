package main

import (
	"fmt"
	"time"
)

func SwapElement(in []int, i int, j int) {
	temp := in[j]
	in[j] = in[i]
	in[i] = temp
}

func BubbleSort(inData []int) {
	fmt.Println("BubbleSort: Begin")
	startTime := time.Now()
	defer func() {
		endTime := time.Now()
		elapseTime := endTime.Sub(startTime)
		fmt.Println("BubbleSort: End: Time = ", elapseTime.Seconds())
		fmt.Println("BubbleSort: OutData: ", inData)
	}()
	fmt.Println("BubbleSort: Data: ", inData)
	for i := 0; i < len(inData); i++ {
		var isSwapped bool
		for j := 0; j < len(inData)-1-i; j++ {
			if inData[j] > inData[j+1] {
				SwapElement(inData, j, j+1)
				isSwapped = true
			}
		}
		fmt.Println("BubbleSort: LastEle: ", inData[len(inData)-1-i])
		if !isSwapped {
			break
		}
	}
}

func InsertionSort(inData []int) {
	fmt.Println("InsertionSort: Begin")
	startTime := time.Now()
	defer func() {
		endTime := time.Now()
		elapseTime := endTime.Sub(startTime)
		fmt.Println("InsertionSort: End: Time = ", elapseTime.Seconds())
		fmt.Println("InsertionSort: OutData: ", inData)
	}()
	fmt.Println("InsertionSort: Data: ", inData)
	for i := 0; i < len(inData); i++ {
		for j := 0; j < i+1; j++ {
			if inData[i] < inData[j] {
				SwapElement(inData, i, j)
			}
		}
		fmt.Println("InsertionSort: Sorted Array: ", inData[:i+1])
	}
}
