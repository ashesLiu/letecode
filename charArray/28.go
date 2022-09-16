package charArray

import(
	"fmt"
)


func getNext(next []int, s string){
	i:=0
	for j:=1;j<len(s);j++{
		if s[i] == s[j]{
			i++
		}else{
			for i>0 && s[i] != s[j]{
				i = next[i-1]
			}
		}
		next[j] = i
	}
}

func strStr(haystack string, needle string) int {
	next := make([]int, len(needle))
	getNext(next,needle)
	fmt.Println(next)
	return -1
}