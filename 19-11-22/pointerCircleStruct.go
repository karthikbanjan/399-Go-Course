package main

import (
	"fmt"
	"math"
)

type circle2 struct {
	x, y, r float64
}

func area2(c *circle2) float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := circle2{0, 0, 5}
	fmt.Println(area2(&c))
}
