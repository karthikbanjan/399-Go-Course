package main

import "fmt"

func main() {
	x := make([]int, 5)
	var y []int

	for i := range x {
		fmt.Scanln(&x[i])
	}

	for i := range x {
		if x[i] >= 1 && x[i] <= 100 {
			y = append(y, x[i])
		}
	}

	fmt.Println("Numbers between 1 and 100:")
	for _, v := range y {
		fmt.Println(v)
	}
}
