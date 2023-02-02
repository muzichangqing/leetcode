package bilibili

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    fh := &ListNode{Next: head}
    fast, slow := fh, fh 
    for i := 0; i <= n; i++ {
        fast = fast.Next
    }
    for fast != nil {
       fast = fast.Next 
       slow = slow.Next
    }
    slow.Next = slow.Next.Next
    return fh.Next
}
