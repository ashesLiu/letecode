package binaryTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type MyStack []*TreeNode

 func (s *MyStack) Push(x *TreeNode){
	 *s = append(*s, x)
 }
 
 func (s *MyStack) Pop() *TreeNode{
	 old := *s
	 x := old[len(old)-1]
	 *s = old[:len(old)-1]
	 return x
 }
func invertTree(root *TreeNode) *TreeNode {
	stack := &MyStack{}
	var prev *TreeNode
	for root!=nil || len(*stack)>0{
		for root!=nil{
			stack.Push(root)
			root = root.Left
		}
		root = stack.Pop()
		if root.Right==nil || root.Right == prev{
			root.Left, root.Right = root.Right, root.Left
			prev = root
            root = nil
		}else{
			stack.Push(root)
			root = root.Right
		}
	}
    return prev
} 