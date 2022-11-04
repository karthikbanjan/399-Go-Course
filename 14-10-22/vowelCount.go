package main

import "fmt"

func main() {
	var s = "presidency"
	var count = 0

	for _, v := range s {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			count = count + 1
			fmt.Print(v)
		}
	}

	fmt.Println(count)
}
