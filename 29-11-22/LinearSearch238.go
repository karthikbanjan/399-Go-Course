package main

import "fmt"

func linearSearch(arr []int, ele int) {
	for i, v := range arr {
		if v == ele {
			fmt.Println("Element found at index:", i)
			return
		}
	}
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

	fmt.Print("Enter the element to be searched:")
	var ele int
	fmt.Scan(&ele)

	linearSearch(arr, ele)
}
