package array

func sortedSquares(nums []int) []int {
	n := len(nums)

	for i := 0; i < n; i++ {
		nums[i] = nums[i] * nums[i]
	}

	// 冒泡
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				swap(nums, i, j)
			}
		}
	}
	return nums
}

func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}
