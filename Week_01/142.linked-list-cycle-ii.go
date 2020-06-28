package main

import "fmt"

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

func NewListNode(value int) *ListNode {
	return &ListNode{Val: value}
}

func detectCycle1(head *ListNode) *ListNode {
	sentry := &ListNode{Next: head}

	fast, slow := sentry, sentry
	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			slow = sentry
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return fast
		}
	}

	return nil
}

func detectCycle2(head *ListNode) *ListNode {
	// 如果直接nil或者下一个节点为nil,那么就肯定没有环了
	if head == nil || head.Next == nil {
		return nil
	}

	// 定义两个节点,一个为快指针,一个慢指针
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		// 快指针每次跳两步
		fast = fast.Next.Next
		// 慢指针每次跳一步
		slow = slow.Next
		// 当 快指针和慢指针相遇了
		if fast == slow {
			// head 节点开始走,fast指针改成每次走一步,相遇的时候 就是环形的起点
			for head != fast {
				head = head.Next
				fast = fast.Next
			}
			return head
		}
	}
	return nil
}

func main() {
	nums := []int{3, 2, 0, -4}
	var head *ListNode
	prev := NewListNode(0)
	for i := 0; i < len(nums); i++ {
		node := NewListNode(nums[i])
		if i == 0 {
			head = node
		}
		prev.Next, prev = node, node
	}
	prev.Next = head.Next
	fmt.Println(detectCycle1(head).Val)
}
