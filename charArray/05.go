package charArray

func replaceSpace(s string) string {
	ans := make([]byte, 0)
	for i := range s {
		if s[i] == ' ' {
			ans = append(ans, '%')
			ans = append(ans, '2')
			ans = append(ans, '0')
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}
