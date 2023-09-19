package main

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	return isEqual(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isEqual(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isEqual(p.Left, q.Left) && isEqual(p.Right, q.Right)
}
