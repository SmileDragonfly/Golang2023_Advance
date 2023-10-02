package main

func main() {
	indata := RandomArray(30, 10000)
	in1 := make([]int, len(indata))
	copy(in1, indata)
	BubbleSort(in1)
	InsertionSort(indata)
}
