package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	resultVals := []int{}
	var base, count, maxCount int
	var traverse func(node *TreeNode)
	var update func(val int)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		update(node.Val)
		traverse(node.Right)
	}
	update = func(val int) {
		if val == base {
			count++
			return
		}
		if count == maxCount {
			resultVals = append(resultVals, base)
		} else if count > maxCount {
			maxCount = count
			resultVals = []int{base}
		}
		count, base = 1, val
	}
	traverse(root)

	return resultVals
}
