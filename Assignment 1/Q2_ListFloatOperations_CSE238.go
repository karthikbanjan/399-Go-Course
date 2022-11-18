package main

import "fmt"

func input(list *[]float64) {
	var num float64

	for {
		fmt.Scan(&num)

		if num == -1 {
			break
		}

		*list = append(*list, num)
	}

}

func sum(list []float64) float64 {
	var s float64

	for _, v := range list {
		s += v
	}

	return s
}

func common(list1 []float64, list2 []float64) []float64 {
	var comMap = make(map[float64]int)

	for _, v := range list1 {
		comMap[v]++
	}

	for _, v := range list2 {
		comMap[v]++
	}

	var comList []float64

	for k, v := range comMap {
		if v > 1 {
			comList = append(comList, k)
		}
	}

	return comList
}

func main() {
	var list1 []float64
	var list2 []float64

	fmt.Println("Enter the numbers for first list(-1 to Exit):")
	input(&list1)
	fmt.Println("Enter the numbers for second list(-1 to Exit):")
	input(&list2)

	fmt.Println("\nList 1:", list1)
	fmt.Println("List 2:", list2)

	if len(list1) != len(list2) {
		fmt.Println("List length not equal!")
	} else {
		fmt.Println("List length equal.")
	}

	if sum(list1) != sum(list2) {
		fmt.Println("List sum not equal!")
	} else {
		fmt.Println("List sum equal.")
	}

	fmt.Println("Common, if any:", common(list1, list2))
}
