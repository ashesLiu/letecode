package hashTable

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int,0)
	twoSearch := func(target, left, right int){
		l,r:=left,right
		for l<r{
			sum := nums[l]+nums[r]
			if sum == target{
				ans = append(ans, []int{-target, nums[l],nums[r]})
				l++
				r--
			}else if sum < target {
				l++
			}else{
				r--
			}
			for l<r && l>left && nums[l]==nums[l-1]{
				l++
			}
			for l<r && r<right && nums[r] == nums[r+1]{
				r--
			}
		}
	}
	for i:=0;i<len(nums);i++{
		if i>0&&nums[i]==nums[i-1]{
			continue
		}
		if nums[i] > 0{
			break
		}
		twoSearch(-nums[i],i+1,len(nums)-1)
	}
	return ans
}