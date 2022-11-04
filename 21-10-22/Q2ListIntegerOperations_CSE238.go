//By Karthik Banjan

package main

func main() {
	l1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	l2 := []int{11, 12, 13, 14, 15, 16, 17, 18, 9, 10}

	if len(l1) == len(l2) {
		println("Both lists are of same length")
	} else {
		println("Both lists are not of same length")
	}

	sum1 := 0
	sum2 := 0

	for _, v := range l1 {
		sum1 += v
	}
	for _, v := range l2 {
		sum2 += v
	}

	if sum1 == sum2 {
		println("The elements in each list sum up to the same value")
	} else {
		println("The elements in each list do not sum up to the same value")
	}

	for _, v := range l1 {
		for _, v2 := range l2 {
			if v == v2 {
				println("There is a value that occur in both lists")
				return
			}
		}
	}
}
