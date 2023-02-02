package bilibili

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var search func(*TreeNode) (*TreeNode, bool)
    search = func (node *TreeNode) (*TreeNode, bool) {
        if node == nil {
            return nil, false
        }
        lrs, lexists := search(node.Left)
        if lrs != nil {
            return lrs, true
        }
        rrs, rexists := search(node.Right)
        if rrs != nil {
            return rrs, true
        }
        mexists := false
        if node.Val == p.Val || node.Val == q.Val {
            mexists = true
        }
        if lexists && rexists || lexists && mexists || mexists && rexists {
            return node, true
        }
        if lexists || rexists || mexists {
            return nil, true
        }
        return nil, false
    }
    rs, _ := search(root)
    return rs
}
