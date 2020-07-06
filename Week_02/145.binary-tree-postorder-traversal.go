package Week_02

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归
/*
func postorderTraversal(root *TreeNode) []int {
    if root == nil { return nil }
    result := append(postorderTraversal(root.Left), postorderTraversal(root.Right)...)
    result = append(result, root.Val)
    return result
}
*/
// 非递归：后序遍历是 左-右-根，可以通过反转 根-右-左 得到，与非递归前序遍历的形式相同
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		// 出栈输出
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Val)
		// 合法子结点入栈，注意入栈顺序
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
	}
	// 逆序
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
