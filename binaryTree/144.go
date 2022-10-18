package binaryTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func preorderTraversal(root *TreeNode) []int {
    stack := make([](*TreeNode), 0)
	arr := make([]int, 0)
	if root==nil{
		return arr
	}
	stack = append(stack, root)
	for len(stack)>0 {
		p := stack[len(stack)-1]
        arr = append(arr, p.Val)
		stack = stack[:len(stack)-1]
		if p.Right!=nil{
			stack = append(stack, p.Right)
		}
		if p.Left!=nil{
			stack = append(stack, p.Left)
		}
	}
	return arr
}