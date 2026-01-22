package main

import "fmt"

func main() {

}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func Inorder(root *Node) {
	if root == nil {
		return
	}
	Inorder(root.Left)
	fmt.Println(root.Value)
	Inorder(root.Right)
}

func Preorder(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.Value)
	Preorder(root.Left)
	Preorder(root.Right)
}

func Postorder(root *Node) {
	if root == nil {
		return
	}
	Postorder(root.Left)
	Postorder(root.Right)
	fmt.Println(root.Value)
}

func LevelOrder(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Println(current.Value)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}
