package Week_02

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
// 递归
/*
func preorder(root *Node) []int {
    if root == nil { return nil }
    result := []int{ root.Val }
    for _, child := range root.Children {
        result = append(result, preorder(child)...)
    }
    return result
}
*/
// 非递归：根-左-右
func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var result []int
	stack := []*Node{root}
	for len(stack) != 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Val)
		// 为了输出顺序 左->右，这里逆序入栈
		for i := len(root.Children) - 1; i >= 0; i-- {
			if root.Children[i] != nil {
				stack = append(stack, root.Children[i])
			}
		}
	}
	return result
}
