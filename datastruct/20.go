package datastruct

type Stack []byte

func (s *Stack) Push(x byte) {
	*s = append(*s, x)
}

func (s *Stack) Pop() (ans byte) {
	length := len(*s)
	ans = (*s)[length-1]
	*s = (*s)[:length-1]
	return
}

func (s *Stack) Top() (ans byte) {
	length := len(*s)
	ans = (*s)[length-1]
	return
}

func isValid(s string) bool {
	var stack Stack
	for _, v := range []byte(s) {
		switch v {
		case '(', '{', '[':
			stack.Push(v)
		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}
			lf := stack.Top()
			if v == ')' && lf == '(' || v == '}' && lf == '{' || v == ']' && lf == '[' {
				stack.Pop()
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
