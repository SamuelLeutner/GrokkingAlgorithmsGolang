package main

import "fmt"

func main() {
	fmt.Println("The position of your target are: ", BinarySearch([]int{1, 2, 4, 7, 9, 10}, 11))
	fmt.Println("The position of your target are: ", BinarySearch([]int{1, 2, 4, 7, 9, 10}, 10))
	fmt.Println("The position of your target are: ", BinarySearch([]int{5, 8, 10, 20}, 8))
	fmt.Println("The position of your target are: ", BinarySearch([]int{3, 6, 7, 9, 10}, 7))
	fmt.Println("The position of your target are: ", BinarySearch([]int{100, 102, 106, 800}, 106))
}

func BinarySearch(arr []int, target int) int {
	// Binary Search are O(log n) = Logarithmic Time. It's the faster time in big O
	// To resolve log i need to ask me "How many 2s do we multiply to get n?"

	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]

		if guess == target {
			return mid
		}
		if guess > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
