package charArray

func strStr(haystack string, needle string) int {
	i, j:=0, 0
	for i<len(haystack){
		pos := i
		for j = 0; j<len(needle);j++{
			if pos==len(haystack)||needle[j]!=haystack[pos]{
				break
			}
			pos++
		}
		if j == len(needle){
			return i
		}
		i++
	}
	return -1
}