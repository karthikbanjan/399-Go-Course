/*
1.  Ask the user to input a sentence which contains at-least 4 words.
Capitalize first letter of  each word and display complete sentence.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter a sentence with at least four words:")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	words := strings.Fields(input)
	if len(words) < 4 {
		fmt.Println("Please input at least four words. Exiting...")
		return
	}
	for i := 0; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}
	fmt.Println(strings.Join(words, " "))
}
