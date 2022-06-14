/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    
    if head.Next == nil {
        return head
    }
    
    var result ListNode
        
    previousNode := &result
        
    currentNode := head

    
    for currentNode != nil && currentNode.Next != nil {
        previousNode.Next = currentNode.Next
        currentNode.Next = previousNode.Next.Next
        
        previousNode.Next.Next = currentNode
        
        previousNode = currentNode
        currentNode = currentNode.Next
    }
    
    return result.Next
}
