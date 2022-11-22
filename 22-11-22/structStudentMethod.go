package main

import "fmt"

type student struct {
	name  string
	age   int
	marks [5]int
}

func (s student) averageMarks() int {
	sum := 0
	for _, v := range s.marks {
		sum += v
	}
	return sum / len(s.marks)
}

func (s student) studentDetails() {
	fmt.Println("Name:", s.name)
	fmt.Println("Age:", s.age)
	fmt.Println("Average Marks:", s.averageMarks())
}

func main() {
	s := student{"Karthik Banjan", 21, [5]int{99, 98, 97, 96, 95}}
	s.studentDetails()
}
