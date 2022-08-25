package hashTable


func isHappy(n int) bool {
	nums := make(map[int64]struct{})
	num = int64(n)
	void := struct{}{}
	for _,ok:=nums[num];num!=1&&!ok;{
		nums[num]=void
		num = nextNum(num)
	}
	if num != 1{
		return false
	}
	return true
}

func nextNum(n int64)int64{
	var ans int64 = 0
	var off int64 = 10
	for n !=0{
		digits := n%off
		ans += digits*digits
		off*=10
	}
	return ans
}