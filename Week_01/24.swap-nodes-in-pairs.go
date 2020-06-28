package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs1(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev, end := dummy, dummy
	for end.Next != nil && end.Next.Next != nil {
		start := prev.Next
		end = end.Next.Next
		prev.Next, end.Next, start.Next = end, start, end.Next
		prev, end = start, start
	}

	return dummy.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p1, p2 := head, head.Next
	p1.Next, p2.Next = swapPairs(p2.Next), p1

	return p2
}
