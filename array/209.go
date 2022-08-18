package array

func minSubArrayLen(target int, nums []int) int {
	l, r, sum, ans := 0, 0, 0, 0
	for ; r < len(nums); r++ {
		sum += nums[r]
		for sum >= target {
			length := r - l + 1
			if ans == 0 || ans > length {
				ans = length
			}
			sum -= nums[l]
			l++
		}
	}
	return ans
}
