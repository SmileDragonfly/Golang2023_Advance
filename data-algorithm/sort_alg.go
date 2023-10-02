package main

import "fmt"

// inData
func BubbleSort(inData []int) {
	fmt.Println("Begin bubble sort")
	fmt.Println("Data: ", inData)
	for i := 0; i < len(inData)-1; i++ {
		var isSwapped bool
		for j := 0; j < len(inData)-1-i; j++ {
			if inData[j] > inData[j+1] {
				temp := inData[j]
				inData[j] = inData[j+1]
				inData[j+1] = temp
				isSwapped = true
			}
		}
		fmt.Println("LastEle: ", inData[len(inData)-1-i])
		if !isSwapped {
			break
		}
	}
}
