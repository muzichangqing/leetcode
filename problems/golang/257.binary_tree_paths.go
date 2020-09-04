package leetcode

import (
	"strconv"
	"strings"
)

// TreeNode a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var paths []string

func binaryTreePaths(root *TreeNode) []string {
	paths = []string{}
	recordPath(root, "")
	return paths
}

func recordPath(node *TreeNode, path string) {
	if node == nil {
		return
	}
	path = path + strconv.Itoa(node.Val) + "->"
	if node.Left == nil && node.Right == nil {
		paths = append(paths, strings.TrimRight(path, "->"))
		return
	}
	recordPath(node.Left, path)
	recordPath(node.Right, path)
}
