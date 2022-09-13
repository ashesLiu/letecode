package charArray

func reverseWords(s string) string {
	str := []byte(s)
	l, r := 0, len(str)-1

	for l < len(str) && str[l] == ' ' {
		l++
	}

	for r >= 0 && str[r] == ' ' {
		r--
	}

	lw := l
	arr := make([]string,0)
	for i := l; i <= r; i++ {
		if str[i] == ' '{
			temp := string(str[lw:i])
			if temp != ""{
                // fmt.Println(temp)
				arr = append(arr, temp)
			}
			lw = i+1
		}
	}
	temp := string(str[lw:r+1])
	if temp != " "{
        // fmt.Println(temp)
		arr = append(arr, temp)
	}	
	var sb strings.Builder
	for i:= len(arr)-1; i>=0;i--{
		sb.WriteString(arr[i])
		if i!=0 {
			sb.WriteString(" ")
		}
	}
	return sb.String()
}
