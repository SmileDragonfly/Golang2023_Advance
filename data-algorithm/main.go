package main

func main() {
	indata := RandomArray(300, 10000)
	in1 := make([]int, len(indata))
	in2 := make([]int, len(indata))
	copy(in1, indata)
	copy(in2, indata)
	BubbleSort(in1)
	InsertionSort(indata)
	SelectionSort(in2)
}
