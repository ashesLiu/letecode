package charArray

func reverseString(s []byte)  {
	l,r:=0,len(s)-1
	for l < r{
	  swap(s,l,r)
	  l++
	  r--
	}
  }
  
  func swap(s []byte, a ,b int ){
	s[a], s[b] = s[b], s[a]
  }