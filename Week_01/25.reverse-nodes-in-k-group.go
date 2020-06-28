package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup1(head *ListNode, k int) *ListNode {
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	return _reverseKGroup(head, n, k)
}

func _reverseKGroup(head *ListNode, n, k int) *ListNode {
	if n < k {
		return head
	}

	var prev *ListNode
	p, rhead := head, head // rhead 记住翻转后的链表尾，后面需要更新尾部指针
	for i := 0; i < k; i++ {
		p.Next, prev, p = prev, p, p.Next // 循环退出时，prev 指向原链表的第 k 个元素也就是翻转后的表头，p 指向第 k+1 个元素
	}
	rhead.Next = _reverseKGroup(p, n-k, k) // 将子链表按规则翻转后拼接到第 k 个元素上

	return prev
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup2(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}

	prev, end := dummy, dummy
	var next *ListNode
	for {
		start := prev.Next                     // start 指向该组的第一个节点
		for i := 0; i < k && end != nil; i++ { // end 指向该组的第 k 个节点
			end = end.Next
		}
		if end == nil { // 该组不足 k 个节点，任务已完成，直接跳出
			break
		}

		end.Next, next = nil, end.Next               // 记录 end 节点的 Next 指针，方便后面翻转后拼接
		prev.Next, start.Next = reverse(start), next // 拼接，翻转后 start 指向了表尾，end 指向了表头
		prev, end = start, start                     // prev 和 end 指针更新
	}

	// 由于 dummy Next 指针在 prev.Next 更新的时候也是按翻转拼接规则更新的，所以 dummy.Next 指向了翻转后的表头
	return dummy.Next
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	p := head
	for p != nil {
		p.Next, prev, p = prev, p, p.Next
	}
	return prev
}
