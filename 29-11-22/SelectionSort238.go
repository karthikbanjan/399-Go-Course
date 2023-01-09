package main

import "fmt"

func selectionSort(arr []int) {
	for i, _ := range arr {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}

	fmt.Println("Sorted array:", arr)
}

func main() {
	fmt.Print("Enter the size of the array:")
	var size int
	fmt.Scan(&size)

	var arr = make([]int, size)
	fmt.Println("Enter the elements of the array:")
	for i, _ := range arr {
		fmt.Scan(&arr[i])
	}

	selectionSort(arr)
}
