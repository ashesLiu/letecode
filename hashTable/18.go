package hashTable

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ans := make([][]int,0)
	for i:=0;i<len(nums);i++{
	  if i > 0 && nums[i] == nums[i-1]{
		continue
	  }
	  for j := i+1;j<len(nums);j++{
		if j > i+1 && nums[j] == nums[j-1]{
		  continue
		}
		l,r := j+1, len(nums)-1
		subTarget := target - nums[i] -nums[j]
		for l<r {
		  sum := nums[l]+nums[r]
		  if sum == subTarget{
			ans = append(ans, []int{nums[i], nums[j], nums[l], nums[r]})
			l++
			r--
		  }else if sum > subTarget{
			r--
		  }else{
			l++
		  }
		  for l<r && l>j+1 && nums[l]==nums[l-1]{
			l++
		  }
		  for l<r && r<len(nums)-1&&nums[r]==nums[r+1]{
			r--
		  }
		}
	  }
	}
	return ans
  }