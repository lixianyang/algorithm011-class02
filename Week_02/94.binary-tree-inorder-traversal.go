package Week_02

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
/*
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }

    result := append(inorderTraversal(root.Left), root.Val)
    result = append(result, inorderTraversal(root.Right)...)

    return result
}
*/
// 非递归
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0, 32)
	stack := make([]*TreeNode, 0, 32)
	// 栈非空且 root 不为空
	for len(stack) != 0 || root != nil {
		// 走到最左路节点，沿途节点都入栈
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 出栈，并输出相应元素
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Val)
		// 进入右子树处理
		root = root.Right
	}
	return result
}
