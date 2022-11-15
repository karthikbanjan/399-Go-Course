package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wordOccurrence(input string) {
	wo := strings.Fields(input)
	wm := make(map[string]int)

	for _, v := range wo {
		wm[v]++
	}

	fmt.Println(wm)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a string: ")
	input, _ := reader.ReadString('\n')

	wordOccurrence(input)
}
