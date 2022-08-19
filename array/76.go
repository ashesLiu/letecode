package array

func minWindow(s string, t string) string {
	hashmap := make(map[byte]int)
	window := make(map[byte]int)
	cnt := 0
	ans := ""
	for i := 0; i < len(t); i++ {
		hashmap[t[i]]++
		cnt++
	}
	i, left := 0, 0
	for i < len(s) {
		window[s[i]]++
		i++
		for check(hashmap, window) {
			if ans == "" || len(ans) > (i-left) {
				ans = s[left:i]
			}
			window[s[left]]--
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
