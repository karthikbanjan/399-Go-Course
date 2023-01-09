/*
Ask the user to input a complete name that contains at least 3 to 5 parts.
Write a go code to make an abbreviation of the name.

Example:
Mohandas Karamchand Gandhi
output: M.K.G.

Brazil, Russia, India, China, and South Africa
output: B.R.I.C.S
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter a name with at least 3 to 5 parts:")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	words := strings.Fields(input)
	if len(words) < 3 {
		fmt.Println("Please input at least 3 parts. Exiting...")
		return
	}

	for _, word := range words {
		if word == "and" {
			fmt.Printf("%c", word[0])
			break
		}
		fmt.Printf("%c.", word[0])
	}

	fmt.Println("\nBy Karthik Banjan")
}
