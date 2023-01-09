package main

import "fmt"

func binarySearch(arr []int, ele int) {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == ele {
			fmt.Println("Element found at index:", mid)
			return
		}

		if arr[mid] < ele {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	fmt.Println("Element not found.")
}

func main() {
	fmt.Print("Enter the size of the array:")
	var size int
	fmt.Scan(&size)

	var arr = make([]int, size)
	fmt.Println("Enter the elements of the array in ascending order:")
	for i, _ := range arr {
		fmt.Scan(&arr[i])
	}

	fmt.Print("Enter the element to be searched:")
	var ele int
	fmt.Scan(&ele)

	binarySearch(arr, ele)
}
