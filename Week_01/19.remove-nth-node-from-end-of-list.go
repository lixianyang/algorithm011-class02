package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	l := 0
	for p := head; p != nil; p = p.Next {
		l++
	}

	dummy := &ListNode{Next: head}
	p := dummy
	for i := 0; i < l-n; i++ {
		p = p.Next
	}
	p.Next = p.Next.Next

	return dummy.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}

	fast, slow := dummy, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast, slow = fast.Next, slow.Next
	}
	slow.Next = slow.Next.Next

	return dummy.Next
}
