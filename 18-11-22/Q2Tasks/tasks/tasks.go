/*
Make a package which has functions of sort in ascending order, descending order,
sum of numbers, sum of square of numbers. Write a go program to call and use this package
*/

package tasks

func SumOfNumbers(numbers ...int) int {
	var sum int
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func SumOfSquares(numbers ...int) int {
	var sum int
	for _, value := range numbers {
		sum += value * value
	}
	return sum
}

func SortAscending(numbers ...int) []int {
	for i, vi := range numbers {
		for j, vj := range numbers {
			if vi < vj {
				numbers[i], numbers[j] = vj, vi
			}
		}
	}

	return numbers
}

func SortDescending(numbers ...int) []int {
	for i, vi := range numbers {
		for j, vj := range numbers {
			if vi > vj {
				numbers[i], numbers[j] = vj, vi
			}
		}
	}

	return numbers
}
