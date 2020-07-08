package Week_03

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归：DFS
/*
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    l, r := maxDepth(root.Left), maxDepth(root.Right)
    if l < r {
        return r+1
    }
    return l+1
}
*/
// BFS
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		for n := len(queue); n > 0; n-- {
			root = queue[0]
			queue = queue[1:]
			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
		}
		depth++
	}
	return depth
}
