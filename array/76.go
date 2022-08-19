package array

func minWindow(s string, t string) string {
	need := make(map[byte]int)
	cnt := 0
	ans := ""
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left := 0
	for i := 0; i < len(s); i++ {
		// cnt 记录覆盖到的有效字符数
		if need[s[i]] > 0 {
			cnt++
		}
		need[s[i]]--
		// fmt.Printf("%v %v %v %v\n", left, i, string(s[left]), string(s[i]))
		// fmt.Printf("%v\n", need)
		if cnt == len(t) {
			// left右移，移过多余字符
			for need[s[left]] < 0 {
				need[s[left]]++
				left++
			}
			// 更新ans
			if len(ans) == 0 || len(ans) >= (i-left+1) {
				ans = s[left : i+1]
			}
			// 收缩使不满足条件
			cnt--
			need[s[left]]++
			left++
		}
	}
	return ans
}

func check(hashmap, window map[byte]int) bool {
	flag := true
	for k, v := range hashmap {
		if num := window[k]; num < v {
			flag = false
			break
		}
	}
	return flag
}
