package array

func minSubArrayLen(target int, nums []int) int {
	l, r, sum, ans := 0, 0, 0, 0
	for r < len(nums) || sum >= target {
		if sum < target {
			sum += nums[r]
			r++
		} else {
			length := r - l
			if ans == 0 || ans > length {
				ans = length
			}
			sum -= nums[l]
			l++
		}
	}
	return ans
}
