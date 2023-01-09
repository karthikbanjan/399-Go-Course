package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

func (head *node) insert(data int) {
	if head.next == nil {
		head.next = &node{data, nil}
		fmt.Println("Insert Success:", data)
	} else {
		head.next.insert(data)
	}
}

func (head *node) delete(data int) {
	if head.next != nil {
		if head.next.data == data {
			head.next = head.next.next
			fmt.Println("Delete Success:", data)
		} else {
			head.next.delete(data)
		}
	}
}

func (head *node) print() {
	fmt.Print(head.data, " ")

	if head.next != nil {
		head.next.print()
	}
}

func (head *node) printReverse() {
	if head.next != nil {
		head.next.printReverse()
	}

	fmt.Print(head.data, " ")
}

func main() {
	var choice int
	var data int

	fmt.Println("Initializing Linked List..")
	fmt.Print("Enter data: ")
	fmt.Scan(&data)
	ll := linkedList{&node{data, nil}}
	fmt.Println("Linked List Initialized")
	fmt.Println()

	for {
		fmt.Println("Linked List Menu")
		fmt.Println("1. Insert")
		fmt.Println("2. Delete")
		fmt.Println("3. Print")
		fmt.Println("4. Print Reverse")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter data: ")
			fmt.Scan(&data)
			ll.head.insert(data)

		case 2:
			fmt.Print("Enter data: ")
			fmt.Scan(&data)
			ll.head.delete(data)

		case 3:
			ll.head.print()
			fmt.Println()

		case 4:
			ll.head.printReverse()
			fmt.Println()

		case 5:
			return

		default:
			fmt.Println("Invalid choice, try again")
		}

		fmt.Println()
	}
}
