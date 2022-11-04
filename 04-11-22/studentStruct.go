//structure storing data of students

package main

import "fmt"

func main() {
	type student struct {
		name    string
		age     string
		rollNo  int
		section string
		cgpa    float32
	}
	var s1 student
	s1.name = "Karthik Banjan"
	s1.age = "21"
	s1.rollNo = 238
	s1.section = "7CSE5"
	s1.cgpa = 9.35
	fmt.Println(s1)
}
