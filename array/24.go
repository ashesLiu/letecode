package array

func removeElement(nums []int, val int) int {
	old, new := 0, 0
	for old < len(nums) {
		if nums[old] != val {
			nums[new] = nums[old]
			new++
		}
		old++
	}
	return new
}
