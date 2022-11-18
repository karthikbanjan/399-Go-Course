package main

import "fmt"

const firstClassTotal, secondClassTotal, economyTotal = 10, 25, 50

func firstClassAvailability(firstClassCurr *int) int {
	return firstClassTotal - *firstClassCurr
}

func secondClassAvailability(secondClassCurr *int) int {
	return secondClassTotal - *secondClassCurr
}

func economyAvailability(economyCurr *int) int {
	return economyTotal - *economyCurr
}

func anyAvailability(firstClassCurr *int, secondClassCurr *int, economyCurr *int) bool {
	return firstClassAvailability(firstClassCurr) > 0 ||
		secondClassAvailability(secondClassCurr) > 0 ||
		economyAvailability(economyCurr) > 0
}

func seatAllotment(class string, num int, cost int, classSelec *int) bool {
	var choice string
	fmt.Println("Do you wish to pay Rs.", cost, "for", num, "tickets in", class, "class? (y/n):")
	fmt.Scanln(&choice)

	if choice != "y" {
		fmt.Println("Restarting")
		return false
	}

	*classSelec += num

	fmt.Println("Seat allotted for", num, "in", class, "class")

	return true
}

func switchCaseClass(class string, num int, firstClassCurr *int, secondClassCurr *int, economyCurr *int) {

	cost := 0

	switch class {
	case "first":
		if firstClassAvailability(firstClassCurr) < num {
			fmt.Println("No seat in this class")
			changeClass(class, num, firstClassCurr, secondClassCurr, economyCurr)
		} else {
			cost = num * 100
			seatAllotment(class, num, cost, firstClassCurr)
		}

	case "second":
		if secondClassAvailability(secondClassCurr) < num {
			fmt.Println("No seat in this class")
			changeClass(class, num, firstClassCurr, secondClassCurr, economyCurr)
		} else {
			cost = num * 50
			seatAllotment(class, num, cost, secondClassCurr)
		}

	case "economy":
		if economyAvailability(economyCurr) < num {
			fmt.Println("No seat in this class")
			changeClass(class, num, firstClassCurr, secondClassCurr, economyCurr)
		} else {
			cost = num * 25
			seatAllotment(class, num, cost, economyCurr)
		}

	default:
		fmt.Println("Invalid class..Restarting")
	}
}

func changeClass(class string, num int, firstClassCurr *int, secondClassCurr *int, economyCurr *int) {
	fmt.Println("Do you wish to change class? (y/n): ")
	var choice string
	fmt.Scanln(&choice)

	if choice == "y" {
		fmt.Println("Enter new class (first, second, economy): ")
		fmt.Scanln(&class)
		switchCaseClass(class, num, firstClassCurr, secondClassCurr, economyCurr)
	}
}

func main() {

	var firstClassCurr, secondClassCurr, economyCurr = 0, 0, 0

	for {
		if !anyAvailability(&firstClassCurr, &secondClassCurr, &economyCurr) {
			fmt.Println("Houseful")
			break
		}

		var class string
		var num int

		fmt.Println("Enter class (first, second, economy): ")
		fmt.Scanln(&class)

		fmt.Println("Enter number of people: ")
		fmt.Scanln(&num)

		switchCaseClass(class, num, &firstClassCurr, &secondClassCurr, &economyCurr)

		fmt.Println("Do you wish to continue? (y/n): ")
		var choice string
		fmt.Scanln(&choice)

		if choice != "y" {
			break
		}
	}
}
