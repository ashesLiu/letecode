package hashTable

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	cnt := 0
	sort.Ints(nums4)
	data1 := make(map[int]int)
	data2 := make(map[int]int)
	for _,vi:=range nums1{
		for _,vj:=nums2{
			data1[vi+vj]++
		}
	}
	for _,vk:=range nums3{
		for _,vl:=nums4{
			data2[vk+vl]++
		}
	}
	for k,v:=range data2{
		cnt += v*data1[k]
	}
	return cnt
	
}