package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	current := root
	for {
		if current.Val > val {
			if current.Left == nil {
				current.Left = &TreeNode{val, nil, nil}
				break
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = &TreeNode{val, nil, nil}
				break
			}
			current = current.Right
		}
	}

	return root
}
