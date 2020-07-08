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
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    result := []int{root.Val}
    result = append(result, preorderTraversal(root.Left)...)
    result = append(result, preorderTraversal(root.Right)...)
    return result
}
*/
// 非递归
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0, 32)
	// 根节点入栈
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		// 出栈并输出
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Val)
		// 符合条件的节点入栈，注意入栈顺序，为了输出先左后右，需要右节点先入
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
	}
	return result
}
// Morris 遍历，O(1) 空间复杂度 参考：https://www.cnblogs.com/anniekim/archive/2013/06/15/morristraversal.html
func preorderTraversal(root *TreeNode) []int {
    var result []int
    var pre *TreeNode
    for root != nil {
        if root.Left == nil {
            result = append(result, root.Val)
            root = root.Right
        } else {
            pre = root.Left
            for pre.Right != nil && pre.Right != root {
                pre = pre.Right
            }
            if pre.Right == nil { // 找到了前驱节点，更新右指针指向当前节点
                pre.Right = root
                result = append(result, root.Val) // 在往左子树走之前，输出当前节点的值
                root = root.Left
            } else {    // pre.Right == root 从前驱节点回退到当前节点
                pre.Right = nil
                root = root.Right
            }
        }
    }
    return result
}
