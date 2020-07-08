package Week_03

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归：根节点到最近叶子节点(叶子节点是指没有子节点的节点)
/*
func minDepth(root *TreeNode) int {
    if root == nil { return 0 }
    if root.Left == nil { return 1 + minDepth(root.Right) }	// 左节点为空，按右节点计算
    if root.Right == nil { return 1 + minDepth(root.Left) }	// 右节点为空，按左节点计算
    l, r := minDepth(root.Left), minDepth(root.Right)	// 左右都不为空，这里选择最短的
    if l < r { return 1 + l }
    return 1 + r
}
*/

// BFS
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]*TreeNode, 0)

	// 起始节点入队
	queue = append(queue, root)

	step := 1
	for len(queue) != 0 {
		n := len(queue) // 记录该层要遍历的节点数
		for i := 0; i < n; i++ {
			// 出队
			p := queue[0]
			queue = queue[1:]

			// 判断是否达到终点
			if p.Left == nil && p.Right == nil {
				return step
			}

			// 把满足条件的子节点加入队列
			if p.Left != nil {
				queue = append(queue, p.Left)
			}
			if p.Right != nil {
				queue = append(queue, p.Right)
			}
		}
		// 增加层数
		step++
	}

	return step
}
