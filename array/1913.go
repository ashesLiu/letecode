package array

func maxProductDifference(nums []int) int {
	l1,l2,r1,r2 := 0,0,0,0
	for i:=range nums{
	  if nums[i]<nums[l1]{
		l1 = i
	  }
	  if nums[i]>nums[r1]{
		r1 = i
	  }
	}
	l2 = r1
	r2 = l1
	for i:=range nums{
	  if nums[i]<nums[l2] && i!= l1{
		l2 = i
	  }
	  if nums[i]>nums[r2] && i!= r1{
		r2 = i
	  }
	}
	return nums[r1]*nums[r2] - nums[l1]*nums[l2]
  }