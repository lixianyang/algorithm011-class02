/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 循环中序遍历
func recoverTree(root *TreeNode)  {
    if root == nil { return }
    var stack []*TreeNode
    prev := &TreeNode{Val: math.MinInt64}
    var p, q *TreeNode
    for len(stack) != 0 || root != nil {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if root.Val < prev.Val {
            if p == nil { 
                p, q = prev, root
            } else { 
                q = root
                break
            }
        }
        prev = root
        root = root.Right
    }
    p.Val, q.Val = q.Val, p.Val
    return
}

// Morris 中序遍历
func recoverTree(root *TreeNode)  {
    var predecessor *TreeNode
    var p, q *TreeNode
    prev := &TreeNode{Val: math.MinInt64}
    for root != nil {
        if root.Left == nil {
            if prev.Val > root.Val {
                q = root
                if p == nil {
                    p = prev
                }
            }
            prev = root
            root = root.Right
        } else {
            predecessor = root.Left
            for predecessor.Right != nil && predecessor.Right != root {
                predecessor = predecessor.Right
            }
            if predecessor.Right == nil {
                predecessor.Right = root
                root = root.Left
            } else {    // predecessor.Right == root
                if prev.Val > root.Val {
                    q = root
                    if p == nil {
                        p = prev
                    }
                }
                predecessor.Right = nil
                prev = root
                root = root.Right
            }
        }
    }
    p.Val, q.Val = q.Val, p.Val
    return
}
