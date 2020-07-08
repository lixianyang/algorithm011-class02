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
// Morris 遍历，利用叶子节点指针存储回退信息，空间复杂度 O(1)
// Morris 详解：https://www.cnblogs.com/anniekim/archive/2013/06/15/morristraversal.html
func inorderTraversal(root *TreeNode) []int {
    var result []int
    var predecessor *TreeNode
    for root != nil {
        if root.Left == nil {
            result = append(result, root.Val)
            root = root.Right
        } else {
            predecessor = root.Left
            for predecessor.Right != nil && predecessor.Right != root {
                predecessor = predecessor.Right
            }
            if predecessor.Right == nil {
                predecessor.Right = root
                root = root.Left    // 如果是前序，需要在往左走前，输出当前节点的值
            } else {    // predecessor.Right == root 表示从前驱节点返回
                result = append(result, root.Val)
                predecessor.Right = nil
                root = root.Right
            }
        }
    }
    return result
}
