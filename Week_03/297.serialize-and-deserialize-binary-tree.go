package Week_03

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var result []string
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		root = queue[0]
		queue = queue[1:]
		if root == nil {
			result = append(result, "")
			continue
		}
		result = append(result, strconv.Itoa(root.Val))
		queue = append(queue, root.Left, root.Right)
	}
	return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	items := strings.Split(data, ",")
	value, _ := strconv.Atoi(items[0])
	root := &TreeNode{Val: value}
	p := root
	queue := []*TreeNode{root}
	for i := 0; len(queue) != 0; {
		p = queue[0]
		queue = queue[1:]
		i++
		if items[i] == "" {
			p.Left = nil
		} else {
			value, _ = strconv.Atoi(items[i])
			p.Left = &TreeNode{Val: value}
			queue = append(queue, p.Left)
		}
		i++
		if items[i] == "" {
			p.Right = nil
		} else {
			value, _ = strconv.Atoi(items[i])
			p.Right = &TreeNode{Val: value}
			queue = append(queue, p.Right)
		}
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
