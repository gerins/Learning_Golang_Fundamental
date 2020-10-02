package main

import "fmt"

func main() {
	tree := &node{data: 5}
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)
	tree.Insert(9)
	tree.Insert(15)

	// fmt.Println(tree.findMin())
	// fmt.Println(tree.findMax())

	tree.preOrderDeepFirstSearch(tree)
}

type node struct {
	left  *node
	data  int
	right *node
}

func (n *node) breathFirstSearch(root *node) {
	var nodeQueue []*node
	nodeQueue = append(nodeQueue, root)

	for nodeQueue != nil {
		fmt.Println(root.data)
		if root.left != nil {
			n.breathFirstSearch(root.left)
		} else if root.right != nil {
			n.breathFirstSearch(root.right)
		}
	}
}

func (n *node) preOrderDeepFirstSearch(root *node) {
	if root == nil {
		return
	}

	fmt.Println(root.data)
	n.preOrderDeepFirstSearch(root.left)
	n.preOrderDeepFirstSearch(root.right)
}

func (n *node) findMin() int {
	if n.left != nil {
		return n.left.findMin()
	}
	return n.data
}

func (n *node) findMax() int {
	if n.right != nil {
		return n.right.findMax()
	}
	return n.data
}

func (n *node) Insert(k int) {
	if k > n.data {
		if n.right == nil {
			n.right = &node{data: k}
		} else {
			n.right.Insert(k)
		}
	} else if k < n.data {
		if n.left == nil {
			n.left = &node{data: k}
		} else {
			n.left.Insert(k)
		}
	}
}

func (n *node) Search(k int) bool {
	if n == nil {
		return false
	}

	if k > n.data {
		return n.right.Search(k)
	} else if k < n.data {
		return n.left.Search(k)
	}

	return true
}
