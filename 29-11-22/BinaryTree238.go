package main

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

type binaryTree struct {
	root *node
}

func (bt *binaryTree) insert(data int) {
	if bt.root == nil {
		bt.root = &node{data, nil, nil}
	} else {
		bt.root.insert(data)
	}
}

func (n *node) insert(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &node{data, nil, nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &node{data, nil, nil}
		} else {
			n.right.insert(data)
		}
	}
}

func (bt *binaryTree) print() {
	if bt.root == nil {
		return
	}

	bt.root.print()
}

func (n *node) print() {
	if n == nil {
		return
	}

	n.left.print()
	fmt.Print(n.data, " ")
	n.right.print()
}

func main() {
	bt := binaryTree{}
	bt.insert(4)
	bt.insert(2)
	bt.insert(1)
	bt.insert(3)
	bt.insert(6)
	bt.insert(5)
	bt.insert(7)
	bt.print()
}
