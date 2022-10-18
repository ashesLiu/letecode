package binaryTree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
//指针的遍历来帮助访问节点，栈则用来处理节点上的元素
func inorderTraversal(root *TreeNode) []int {
	if root == nil{
		return nil
	}
	stack := []*TreeNode{}
	p := root
	ans := make([]int, 0)
	for p!=nil || len(stack)>0{
		if p!=nil{
			stack = append(stack, p)
			p = p.Left
		}else{
			// 栈首元素的左子树一定已经处理完成
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, p.Val)
			p = p.Right
		}
	}
	return ans
}