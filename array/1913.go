package array

func maxProductDifference(nums []int) int {
	var mx1,mx2,mn1,mn2 int = 0,0,1e5,1e5
	for _,v:=range nums{
	  if v>mx1{
		mx2 = mx1
		mx1 = v
	  }else if v>mx2{
		mx2 = v
	  }
  
	  if v< mn1{
		mn2 = mn1
		mn1 = v
	  }else if v<mn2{
		mn2 = v
	  }
	  // fmt.Printf("%v %v %v %v\n",mx1,mx2,mn1,mn2)
	}
	// fmt.Printf("%v %v %v %v\n",mx1,mx2,mn1,mn2)
	return mx1*mx2 - mn1*mn2
}