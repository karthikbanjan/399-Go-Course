// Code to convert temperature from Fahrenheit to Celsius by Karthik Banjan

package main

import "fmt"

func main() {
	var temperatureF float64

	fmt.Print("Enter the temperature in Fahrenheit: ")
	fmt.Scanf("%f", &temperatureF)

	temperatureC := (temperatureF - 32) * 5 / 9

	fmt.Println("Temperature in Celsius is: ", temperatureC)
}
