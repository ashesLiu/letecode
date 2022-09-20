package charArray

import(
	"fmt"
)


func getNext(next []int, s string){
	i:=0
	for j:=1;j<len(s);j++{
		for i>0 && s[i]!=s[j]{
			i = next[i-1]
		}
		if s[i]==s[j]{
			i++
		}
		next[j] = i
	}
}

func strStr(haystack string, needle string) int {
	next := make([]int, len(needle))
	getNext(next,needle)
	j:=0
	for i:=0;i<len(haystack);i++{
		for j>0 && haystack[i]!=needle[j]{
			j = next[j-1]
		}
		if haystack[i]==needle[j]{
			j++
			if j==len(needle){
				return i - j + 1
			}
		}
	}
	return -1
}