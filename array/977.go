package array

import (
	"math/rand"
)

func SortedSquares(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		nums[i] = nums[i] * nums[i]
	}

	// 快排
	QuickSort(nums, 0, len(nums))
	return nums
}

// quichSort: [left,right)
func QuickSort(nums []int, left, right int) {
	if left < right {
		idx := partition2(nums, left, right)
		QuickSort(nums, left, idx)
		QuickSort(nums, idx+1, right)
	}
}

// partition: 单路
func partition(nums []int, l, r int) int {

	// 选择任意值为枢轴元素
	pivot := rand.Intn(r-l) + l
	swap(nums, pivot, l)
	v := nums[l]
	// 双指针: i遍历 pos为已遍历元素的划分点
	// 循环不变量： nums[l:pos)<pivot nums[pos:i)>=pivot
	pos := l
	for i := l + 1; i < r; i++ {
		if nums[i] < v {
			pos += 1
			swap(nums, pos, i)
		}
	}
	swap(nums, pos, l)
	return pos
}

// partition2: 双路
func partition2(nums []int, left, right int) int {

	// 选择任意值为枢轴元素
	pivot := rand.Intn(right-left) + left
	swap(nums, pivot, left)
	v := nums[left]
	// 双指针 l r
	// 循环不变量 nums[left+1,l]<=v nums[r,right)>=v
	l, r := left+1, right-1
	for l <= r {
		for l <= r && nums[l] < v {
			l++
		}
		for l <= r && nums[r] > v {
			r--
		}
		if l < r {
			swap(nums, l, r)
			l++
			r--
		} else {
			break
		}
	}
	swap(nums, r, left)
	return r
}

func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}
