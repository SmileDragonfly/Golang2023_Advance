package main

func main() {
	indata := RandomArray(500, 1000)
	in1 := make([]int, len(indata))
	in2 := make([]int, len(indata))
	in3 := make([]int, len(indata))
	in4 := make([]int, len(indata))
	copy(in1, indata)
	copy(in2, indata)
	copy(in3, indata)
	copy(in4, indata)
	BubbleSort(in1)
	InsertionSort(indata)
	SelectionSort(in2)
	MergerSort(in3)
	QuickSort(in4)
}
