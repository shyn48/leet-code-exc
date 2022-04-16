/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    resultList := &ListNode{
        Val: 0,
        Next: nil,
    }
    
    temp := resultList
        
    carry := 0
    counter := 0
    for l1 != nil || l2 != nil {
        counter++        
        leftOperand := 0
        if l1 != nil {
            leftOperand = l1.Val
        }
        
        rightOperand := 0
        if l2 != nil {
            rightOperand = l2.Val
        }
                
        sum := leftOperand + rightOperand + carry

        carry = sum / 10
        sum = sum % 10

        newNode := &ListNode{
            Val: sum,
            Next: nil,
        }
    
        temp.Next = newNode
        temp = temp.Next
        
        if l1 != nil {
           l1 = l1.Next
        }
        if l2 != nil {
           l2 = l2.Next
        }
    }
    
    if carry > 0 {
        temp.Next = &ListNode{
            Val: carry,
            Next: nil,
        }
    }
    
    resultList = resultList.Next
                                   
    return resultList
}
