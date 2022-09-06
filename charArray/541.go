package charArray

func reverseStr(s string, k int) string {
	length := len(s)
	str := []rune(s)
	for i := 0; i < length; i += 2 * k {
		if length-i < k {
			reverse(str[i:length])
		} else if length-i < 2*k {
			reverse(str[i : i+k])
		} else {
			reverse(str[i : i+k])
		}
	}
	return string(str)
}

func reverse(s []rune) {
	//fmt.Printf("reverse addr:%p", &s)
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
