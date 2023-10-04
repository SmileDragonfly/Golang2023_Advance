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
	//fmt.Println("BubbleSort: Begin")
	startTime := time.Now()
	defer func() {
		endTime := time.Now()
		elapseTime := endTime.Sub(startTime)
		fmt.Println("BubbleSort: End: Time = ", elapseTime.Seconds())
		//fmt.Println("BubbleSort: OutData: ", inData)
	}()
	//fmt.Println("BubbleSort: Data: ", inData)
	for i := 0; i < len(inData); i++ {
		var isSwapped bool
		for j := 0; j < len(inData)-1-i; j++ {
			if inData[j] > inData[j+1] {
				SwapElement(inData, j, j+1)
				isSwapped = true
			}
		}
		//fmt.Println("BubbleSort: LastEle: ", inData[len(inData)-1-i])
		if !isSwapped {
			break
		}
	}
}

func InsertionSort(inData []int) {
	//fmt.Println("InsertionSort: Begin")
	startTime := time.Now()
	defer func() {
		endTime := time.Now()
		elapseTime := endTime.Sub(startTime)
		fmt.Println("InsertionSort: End: Time = ", elapseTime.Seconds())
		//fmt.Println("InsertionSort: OutData: ", inData)
	}()
	//fmt.Println("InsertionSort: Data: ", inData)
	for i := 0; i < len(inData); i++ {
		for j := 0; j < i+1; j++ {
			if inData[i] < inData[j] {
				SwapElement(inData, i, j)
			}
		}
		//fmt.Println("InsertionSort: Sorted Array: ", inData[:i+1])
	}
}

func SelectionSort(in []int) {
	//fmt.Println("SelectionSort: Begin")
	startTime := time.Now()
	defer func() {
		endTime := time.Now()
		elapseTime := endTime.Sub(startTime)
		fmt.Println("SelectionSort: End: Time = ", elapseTime.Seconds())
		//fmt.Println("SelectionSort: OutData: ", in)
	}()
	for i := 0; i < len(in); i++ {
		for j := i + 1; j < len(in); j++ {
			minIndex := i
			if in[j] < in[minIndex] {
				minIndex = j
			}
			SwapElement(in, i, minIndex)
		}
	}
}

func MergerSort(in []int) {
	//fmt.Println("MergerSort: Begin")
	startTime := time.Now()
	defer func() {
		endTime := time.Now()
		elapseTime := endTime.Sub(startTime)
		fmt.Println("MergerSort: End: Time = ", elapseTime.Seconds())
		//fmt.Println("MergerSort: OutData: ", in)
	}()
	in = MergerSortRecursive(in)
}

func MergerSortRecursive(l []int) []int {
	if len(l) == 1 {
		return l
	}
	middle := len(l) / 2
	l1 := l[:middle]
	l2 := l[middle:]
	l1 = MergerSortRecursive(l1)
	l2 = MergerSortRecursive(l2)
	return MergerList(l1, l2)
}

func MergerList(l1 []int, l2 []int) []int {
	ret := make([]int, 0)
	for len(l1) > 0 && len(l2) > 0 {
		if l1[0] < l2[0] {
			ret = append(ret, l1[0])
			l1 = l1[1:]
		} else {
			ret = append(ret, l2[0])
			l2 = l2[1:]
		}
	}
	ret = append(ret, l1...)
	ret = append(ret, l2...)
	return ret
}

func QuickSort(l []int) {
	startTime := time.Now()
	defer func() {
		endTime := time.Now()
		elapseTime := endTime.Sub(startTime)
		fmt.Println("QuickSort: End: Time = ", elapseTime.Seconds())
		fmt.Println("QuickSort: OutData: ", l)
	}()
	QuickSortRecursive(l)
}

func QuickSortRecursive(l []int) {
	if len(l) <= 1 {
		return
	}
	pivot := l[len(l)-1]
	i := 0
	j := len(l) - 2
	for {
		for l[i] < pivot {
			i++
		}
		for l[j] > pivot && j > 0 {
			j--
		}
		if i >= j {
			break
		}
		if l[i] > l[j] {
			SwapElement(l, i, j)
			i++
			j--
		}
	}
	SwapElement(l, i, len(l)-1)
	QuickSortRecursive(l[0:i])
	QuickSortRecursive(l[i:len(l)])
}
