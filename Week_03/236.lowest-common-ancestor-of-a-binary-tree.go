package Week_03

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
// 详细题解：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/solution/236-er-cha-shu-de-zui-jin-gong-gong-zu-xian-hou-xu/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left, right := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
	// 左右子树分别有一个
	if left != nil && right != nil {
		return root
	}
	// p, q 都在左子树
	if left != nil {
		return left
	}
	// p, q 都在右子树
	if right != nil {
		return right
	}
	return nil
}

// 暴力：分别求出p, q 的所有祖先，在遍历比较
/*
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    pa := Ancestors(root, p)
    qa := Ancestors(root, q)
    l1, l2 := len(pa), len(qa)
    l := l1
    if l > l2 {
        l = l2
    }
    var pre *TreeNode
    for i := 0; i < l; i++ {
        if pa[i].Val != qa[i].Val {
            break
        }
        pre = pa[i]
    }
    return pre
}

func Ancestors(root *TreeNode, p *TreeNode) []*TreeNode {
    if p.Val == root.Val {
        return []*TreeNode{ root }
    }
    if root.Left != nil {
        if l := Ancestors(root.Left, p); l != nil {
            return append([]*TreeNode{ root }, l...)
        }
    }
    if root.Right != nil {
        if r := Ancestors(root.Right, p); r != nil {
            return append([]*TreeNode{ root }, r...)
        }
    }
    return nil
}
*/
