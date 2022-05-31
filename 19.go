/**
* Definition for singly-linked list.
// * type ListNode struct {
// *     Val int
// *     Next *ListNode
// * }
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var ptr *ListNode
	lenList := 0

	ptr = head

	for ptr != nil {
		ptr = ptr.Next
		lenList++
	}

	if lenList == 0 {
		return head
	}

	if lenList == n {
		return head.Next
	}

	ptr = head
	currIndex := 0

	var nodeBeforeNthNode *ListNode
	for ptr != nil {
		if currIndex+1 == lenList-n {
			nodeBeforeNthNode = ptr
		}

		if currIndex == lenList-n {
			nodeBeforeNthNode.Next = ptr.Next
		}
		ptr = ptr.Next
		currIndex++
	}

	return head
}