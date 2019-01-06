package goleet

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	// create first linkedlist
	var l1 ListNode
	var l1Next ListNode
	var l1NextNext ListNode
	l1.Val = 1
	l1Next.Val = 2
	l1Next.Val = 4
	l1.Next = &l1Next
	l1Next.Next = &l1NextNext
	l1NextNext.Next = nil
	// second linkedlist
	var l2 ListNode
	var l2Next ListNode
	var l2NextNext ListNode
	l2.Val = 1
	l2Next.Val = 3
	l2Next.Val = 4
	l2.Next = &l2Next
	l2Next.Next = &l2NextNext
	l2NextNext.Next = nil
	// create expected linkedlist
	var expected ListNode
	var expectedPtr *ListNode
	var exp1 ListNode
	var exp2 ListNode
	var exp3 ListNode
	var exp4 ListNode
	var exp5 ListNode
	expected.Val = 1
	exp1.Val = 1
	exp2.Val = 2
	exp3.Val = 3
	exp4.Val = 4
	exp5.Val = 4
	expectedPtr = &expected
	expected.Next = &exp1
	exp1.Next = &exp2
	exp2.Next = &exp3
	exp3.Next = &exp4
	exp4.Next = &exp5
	exp5.Next = nil

	numCases := 1
	for i := 0; i < numCases; i++ {
		actual := mergeTwoLists(&l1, &l2)

		for {
			if actual == nil || expectedPtr == nil {
				break
			}
			if actual.Val != expectedPtr.Val {
				t.Fatalf("fail")
			}
			actual = actual.Next
			expectedPtr = expectedPtr.Next
		}
		fmt.Println("case ", i, " PASS")
	}
}
