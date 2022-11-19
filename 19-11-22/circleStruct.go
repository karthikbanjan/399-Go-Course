package main

import (
	"fmt"
	"math"
)

type circle struct {
	x, y, r float64
}

func area(c circle) float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := circle{0, 0, 5}
	fmt.Println(area(c))
}
