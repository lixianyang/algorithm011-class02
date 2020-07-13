/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil { return nil }
    result := make([][]int, 0)
    queue := []*TreeNode{ root }
    for len(queue) > 0 {
        n := len(queue)
        level := make([]int, n)
        for i := 0; i < n; i++ {
            root = queue[0]
            queue = queue[1:]
            level[i] = root.Val
            if root.Left != nil { queue = append(queue, root.Left) }
            if root.Right != nil { queue = append(queue, root.Right) }
        }
        result = append(result, level)
    }
    return result
}
