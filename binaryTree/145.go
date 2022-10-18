package binaryTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func postorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil{
		return ans
	}
	stack := []*TreeNode{root}
	for len(stack)>0{
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
        ans = append(ans, p.Val)
		if p.Left!=nil{
			stack = append(stack, p.Left)
		}
		if p.Right!=nil{
			stack = append(stack, p.Right)
		}
	}
	reverse(ans)
	return ans
}

func reverse(arr []int){
	l,r := 0, len(arr)-1
	for l<r{
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}