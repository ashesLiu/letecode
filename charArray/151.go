package charArray

func reverseWords(s string) string {
	//reverse [left, right)
	reverse := func(data []byte, left, right int) {
		l, r := left, right-1
		for l < r {
			data[l], data[r] = data[r], data[l]
			l++
			r--
		}
	}
	str := []byte(s)
	l, r := 0, len(str)-1
	// 去掉前后空格
	for l < len(str) && str[l] == ' ' {
		l++
	}
	for r >= 0 && str[r] == ' ' {
		r--
	}
	low, fast := l, l
	// 去掉中间多余空格
	for fast <= r {
		if !(str[fast] == ' ' && fast-1 > 0 && str[fast] == str[fast-1]) {
			str[low] = str[fast]
			low++
		}
		fast++
	}
	str = str[l:low]
	// 反转整个str
	reverse(str, 0, len(str))

	// 反转单词
	low, fast = 0, 0
	for ;fast <= len(str);fast++ {
		if fast == len(str)||str[fast] == ' ' {
			reverse(str, low, fast)
			low = fast + 1
		}
	}
	return string(str)
}