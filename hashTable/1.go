package hashTable

func twoSum(nums []int, target int) []int {
	data := make(map[int]int)
	for i,v := range nums{
		if j,ok:= data[target-v];ok{
			return []int{j,i}
		}
		data[v] = i
	}
	return []int{-1,-1}
}