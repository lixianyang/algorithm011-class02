/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 最大值在 root.Val, root.Val+left, root.Val+right, root.Val+left+right 处取得
// 返回上一级时，只能返回 root.Val, root.Val+left, root.Val+right 中最大的，不能返回 root.Val+left+right
func maxPathSum(root *TreeNode) int {
    max := math.MinInt64
    doMaxPathSum(root, &max)
    return max
}
func doMaxPathSum(root *TreeNode, max *int) int {
    if root == nil { return 0 }
    left := maxValue(0, doMaxPathSum(root.Left, max))
    right := maxValue(0, doMaxPathSum(root.Right, max))
    *max = maxValue(*max, left+right+root.Val)
    return root.Val + maxValue(left, right)
}
func maxValue(a, b int) int {
    if b > a { return b }
    return a
}
