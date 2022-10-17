package datastruct

func removeDuplicates(s string) string {
	var stack []byte
	for _, v := range []byte(s) {
		if len(stack) > 0 && stack[len(stack)-1] == v {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return string(stack)
}
