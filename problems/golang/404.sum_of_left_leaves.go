package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	var traverse func(node *TreeNode, isLeft bool)
	traverse = func(node *TreeNode, isLeft bool) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil && isLeft {
			sum += node.Val
			return
		}
		traverse(node.Left, true)
		traverse(node.Right, false)
	}
	traverse(root, false)
	return sum
}
