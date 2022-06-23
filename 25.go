/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(begin *ListNode, end *ListNode) *ListNode{
    curr := begin.Next
    prev := begin
    first := curr
    
    for curr != end {
        recordNext := curr.Next
        curr.Next = prev
        prev = curr
        curr = recordNext
    }
    begin.Next = prev
    first.Next = curr
    return first
}


func reverseKGroup(head *ListNode, k int) *ListNode {
    var begin *ListNode
    if head == nil || head.Next == nil || k == 1 {
        return head
    }
    
    dummy := &ListNode{
        Val: 0,
        Next: head,
    }
    
    begin = dummy  
    
    i := 0
    for head != nil {
        i++
        if i%k == 0 {
            begin = reverseBetween(begin, head.Next)
            head = begin.Next
        } else {
            head = head.Next
        }
    }
    
    return dummy.Next
}
