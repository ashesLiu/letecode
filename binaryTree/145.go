package binaryTree

type Stack []*TreeNode

func (s *Stack) Push(x *TreeNode){
	*s = append(*s, x)
}

func (s *Stack) Pop() *TreeNode{
	old := *s
	x := old[len(old)-1]
	*s = old[:len(old)-1]
	return x
}

func (s *Stack) Len() int{
	return len(*s)
}

func postorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	p := root
	stack := &Stack{}
	var prev *TreeNode
	for p!=nil || stack.Len()>0{
		for p!=nil{
			stack.Push(p)
			p = p.Left
		}
		p = stack.Pop()
		if p.Right == nil || p.Right == prev{
			ans = append(ans, p.Val)
			prev = p
            p = nil
		}else{
			stack.Push(p)
			p = p.Right
		}
	}
	return ans
}