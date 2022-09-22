package charArray

func repeatedSubstringPattern(s string) bool {
	for i:=len(s)/2;i>=1;{
		subStr := s[:i]
		k,j :=0, i
        // fmt.Println(subStr)
		for ;j<len(s);j++{
            // fmt.Printf("%c %c\n",subStr[k], s[j])
			if subStr[k] == s[j]{
                if j==len(s)-1 && k==i-1{
                    return true
                }
                k++
				if k==i{
					k=0
				}
			}else{
                break
			}
		}
        i--
	}
	return false
}
