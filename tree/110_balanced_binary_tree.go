package main

import "math"

func isBalanced(root *TreeNode) bool {
	balanced, _ := dfs(root)
	return balanced
}

func dfs(root *TreeNode) (isBalanced bool, height int) {
	if root == nil {
		return true, 0
	}

	left, leftHeight := dfs(root.Left)
	right, rightHeight := dfs(root.Right)

	balanced := left && right && math.Abs(float64(leftHeight)-float64(rightHeight)) <= 1

	return balanced, 1 + int(math.Max(float64(leftHeight), float64(rightHeight)))
}
