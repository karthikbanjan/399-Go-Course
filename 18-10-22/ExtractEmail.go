package main

import "fmt"

func isValidEmail(email string) bool {
	dotPos := -1
	atPos := -1

	for i, v := range email {
		if v == '.' {
			dotPos = i
		} else if v == '@' {
			atPos = i
		}
	}

	if dotPos != -1 && atPos != -1 && dotPos > atPos {
		return true
	}

	return false
}

func main() {
	email := []string{"karthik", "karthik@gmail.com", "sam@gmail.com"}

	for _, v := range email {
		if isValidEmail(v) {
			fmt.Println(v)
		}
	}
}
