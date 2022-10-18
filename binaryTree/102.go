package binaryTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	if root != nil{
		queue = append(queue, root)
	}
	ans := make([][]int, 0)
	for len(queue)>0{
		size := len(queue)
		tmp := make([]int, size)
		for i:=0;i<size;i++{
			tmp[i] = queue[i].Val
			if queue[i].Left!=nil{
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right!=nil{
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[size:]
		ans = append(ans, tmp)
	}
	return ans
}