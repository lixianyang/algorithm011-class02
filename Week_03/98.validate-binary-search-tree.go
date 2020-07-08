package Week_03

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 树值是否在指定区间
/*
func isValidBST(root *TreeNode) bool {
    return isValid(root, math.MinInt64, math.MaxInt64)
}
func isValid(root *TreeNode, lower, upper int) bool {
    if root == nil { return true }
    if root.Val >= upper || root.Val <= lower { return false }
    return isValid(root.Left, lower, root.Val) && isValid(root.Right, root.Val, upper)
}
*/
// 二叉树中序遍历有序
func isValidBST(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	pre := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= pre {
			return false
		}
		pre = root.Val
		root = root.Right
	}
	return true
}
