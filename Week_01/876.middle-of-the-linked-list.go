package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode1(head *ListNode) *ListNode {
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}

	n = n/2 + 1
	for i := 1; i < n; i++ {
		head = head.Next
	}

	return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode2(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy
	for fast.Next != nil && fast.Next.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
	}
	return slow.Next
}
