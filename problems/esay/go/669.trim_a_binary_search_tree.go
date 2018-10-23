package leetcode

func trimBST(root *TreeNode, L int, R int) *TreeNode {
    if root == nil {
        return root
    }
    if root.Val < L {
        return trimBST(root.Right, L, R)
    }
    if root.Val > R {
        return trimBST(root.Left, L, R)
    }
    root.Left = trimBST(root.Left, L, R)
    root.Right = trimBST(root.Right, L, R)
    return root
}
