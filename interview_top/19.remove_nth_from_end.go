package interviewtop

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var (
        fakerHead ListNode
        nodeStack []*ListNode
    )
    fakerHead.Next = head
    head = &fakerHead
    for head != nil {
        nodeStack = append(nodeStack, head)
        head = head.Next
    }
    removedIndex := len(nodeStack) - n  
    nodeStack[removedIndex-1].Next = nodeStack[removedIndex].Next
    return nodeStack[0].Next
}
