package hashTable

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int,0)
	for i:=0;i<len(nums);i++{
		if i>0&&nums[i]==nums[i-1]{
			continue
		}
		for j:=i+1;j<len(nums);j++{
			if i>0&&nums[i]==nums[i-1]{
				continue
			}
			if BinarySearch(nums,-nums[i]-nums[j],j+1,len(nums)-1){
				ans = append(ans, []int{nums[i],nums[j],-nums[i]-nums[j]})	
			}
		}
	}
	return ans
}

func BinarySearch(nums []int,target,left,right int) bool{
	l,r:=left,right
	for l<=r{
		mid := l + (r-l)/2
		if nums[mid] == target{
			return true
		}else if nums[mid]<target{
			l = mid+1
		}else{
			r = mid-1
		}
	}
	return false
}