package charArray

import(
	"fmt"
)

func strStr(haystack string, needle string) int {
	next := make([]int, len(needle))
	i:=0
	for j:=1;j<len(needle);j++{
		pi, pj := i, j
		for pi<pj{
			if needle[pi] == needle[pj]{
				pi++
				pj--
				next[j]++
			}else{
				break
			}
		}
	}
	fmt.Println(next)
	return -1
}