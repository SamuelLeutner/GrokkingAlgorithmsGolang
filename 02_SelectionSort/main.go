package main

import "fmt"

func main() {
	fmt.Println("The index of the smallest number in this array, are: ", findSmallest([]int{4, 6, 7, 1}))
	fmt.Println("The selection sort is: ", SelectionSort([]int{4, 6, 7, 1}))
	fmt.Println("The selection sort is: ", SelectionSort([]int{2, 1, 6, 8, 10, 90}))
	fmt.Println("The selection sort is: ", SelectionSort([]int{9, 1, 4, 56, 7}))
}

// O(n)
func findSmallest(arr []int) int {
	s := arr[0]
	sIndx := 0

	for i, v := range arr {
		if v < s {
			s = v
			sIndx = i
		}
	}

	return sIndx
}

// O(n^2)
func SelectionSort(arr []int) []int {
	// O(1)
	newArr := make([]int, len(arr))

	for i := range arr {
		// O(n)
		sIndx := findSmallest(arr)
		newArr[i] = arr[sIndx]
		arr = append(arr[:sIndx], arr[sIndx+1:]...)
	}

	return newArr
}
