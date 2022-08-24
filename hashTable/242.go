package hashTable

func isAnagram(s string, t string) bool {
    cnt := make(map[rune]int)
	if len(s)!=len(t){
		return false
	}
	for _,v := range s{
		cnt[v]++
	}
	for _,v := range t{
		cnt[v]--
		if cnt[v]<0{
			return false
		}
	}
	return true
}