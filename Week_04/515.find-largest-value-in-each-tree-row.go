/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
    if root == nil { return nil }
    result := make([]int, 0)
    queue := []*TreeNode{ root }
    for len(queue) > 0 {
        max := queue[0].Val
        n := len(queue)
        for i := 0; i < n; i++ {
            root = queue[0]
            queue = queue[1:]
            if root.Val > max {
                max = root.Val
            }
            if root.Left != nil { queue = append(queue, root.Left) }
            if root.Right != nil { queue = append(queue, root.Right) }
        }
        result = append(result, max)
    }
    return result
}
