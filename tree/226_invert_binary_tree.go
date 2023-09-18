package main

func invertTree(root *TreeNode) *TreeNode {
	traverse(root)
	return root
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	left := root.Left
	right := root.Right

	root.Left = right
	root.Right = left

	traverse(root.Left)
	traverse(root.Right)
}
