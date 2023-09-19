package main

import (
	"math"
)

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	dfsRunner(root, &diameter)
	return diameter

}

func dfsRunner(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}

	leftHeight := dfsRunner(root.Left, diameter)
	rightHeight := dfsRunner(root.Right, diameter)

	currentNodeHeight := 1 + int(math.Max(float64(leftHeight), float64(rightHeight)))

	if *diameter < leftHeight+rightHeight {
		*diameter = leftHeight + rightHeight
	}

	return currentNodeHeight
}
