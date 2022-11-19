package main

import "math"

type circle3 struct {
	x, y, r float64
}

func (c circle3) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := circle3{0, 0, 5}
	println(c.area())
}
