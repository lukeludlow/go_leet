package goleet

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var answer *ListNode
	dummy := answer
	for {
		if l1 == nil || l2 == nil {
			break
		}
		if l1.Val < l2.Val {
			answer.Next = l1
			answer.Next.Next = l2
			l1 = l1.Next
			l2 = l2.Next
		} else {
			answer.Next = l2
			answer.Next.Next = l1
			l1 = l1.Next
			l2 = l2.Next
		}
		answer = answer.Next.Next
	}
	return dummy.Next
}
