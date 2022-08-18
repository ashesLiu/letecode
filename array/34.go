package array

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] == target {
			i, j := mid-1, mid+1
			for i >= 0 && nums[i] == target {
				i--
			}
			for j < len(nums) && nums[j] == target {
				j++
			}
			return []int{i + 1, j - 1}
		} else {
			r = mid - 1
		}
	}
	return []int{-1, -1}
}
