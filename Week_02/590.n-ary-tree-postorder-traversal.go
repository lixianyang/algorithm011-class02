package Week_02

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
type Node struct {
	Val      int
	Children []*Node
}

// 递归：左-右-根
/*
func postorder(root *Node) []int {
    if root == nil { return nil }
    var result []int
    for _, child := range root.Children {
        result = append(result, postorder(child)...)
    }
    result = append(result, root.Val)
    return result
}
*/
// 非递归：左-右-根 可以先得到 根-右-左 再反转
func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var result []int
	stack := []*Node{root}
	for len(stack) != 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Val)
		// 为了使输出 右->左，这里从 左->右 进栈
		for _, child := range root.Children {
			if child != nil {
				stack = append(stack, child)
			}
		}
	}
	// 反转
	for l, r := 0, len(result)-1; l < r; l, r = l+1, r-1 {
		result[l], result[r] = result[r], result[l]
	}
	return result
}
