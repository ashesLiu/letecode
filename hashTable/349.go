package hashTable

func intersection(nums1 []int, nums2 []int) []int {
	data := make(map[int]int)
	ans := make([]int, 0)
	for _,v := range nums1{
		if _,ok:=data[v];!ok{
			data[v]++
		}
	}
	for _,v := range nums2{
		if _,ok:=data[v];ok{
			data[v]--
		}
	}
    // fmt.Println(data)
	for i,v := range data{
		if v <= 0{
			ans = append(ans,i)
		}
	}
	return ans
}