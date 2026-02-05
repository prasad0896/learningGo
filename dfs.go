package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return 1 + max(leftDepth, rightDepth)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	
	fmt.Println("Tree maximum depth is: ", maxDepth(tree))
}

