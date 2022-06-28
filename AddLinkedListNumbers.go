package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resultList := &ListNode{
		Val:  0,
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
			Val:  sum,
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
			Val:  carry,
			Next: nil,
		}
	}

	resultList = resultList.Next

	return resultList
}

// second run
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    result := &ListNode{
        Val: 0,
        Next: nil,
    }
    
    dummy := result
    carry := 0
    for l1 != nil || l2 != nil {
        x,y := 0, 0
        fmt.Println(l1,l2,dummy)
        if l1 != nil {
            x = l1.Val
        }
        if l2 != nil {
            y = l2.Val
        }

        sum := x + y + carry
        carry = sum / 10
        dummy.Next = &ListNode{
            Val: sum % 10,
        }

        
        dummy = dummy.Next        
        if l1 != nil {
            l1 = l1.Next
        }
        if l2 != nil {
            l2 = l2.Next
        }
    }
    
    if carry > 0 {
        dummy.Next = &ListNode{
            Val: carry,
        }
    }
    
    return result.Next
}
