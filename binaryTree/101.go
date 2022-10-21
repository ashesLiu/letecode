package binaryTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetric(root *TreeNode) bool {
	return inner(root, root)
}

func inner(left, right *TreeNode) bool{
	if left == nil && right == nil{
		return true
	}else if left == nil || right == nil{
		return false
	}else if inner(left.Left, right.Right){
        if left.Val != right.Val{
			return false
		}
		return inner(left.Right, right.Left)	
	}
    return false
}


