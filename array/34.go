package array

import "fmt"

func searchRange(nums []int, target int) []int {
	return []int{binarySearchLeft(nums, target), binarySearchRight(nums, target)}
}

func binarySearchLeft(nums []int, target int) int {
	l, r := 0, len(nums)-1
	leftBorder := -1
	for l <= r {
		mid := l + (r-l)/2
		fmt.Printf("%v %v %v\n", l, r, mid)
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] == target {
			r = mid - 1
			leftBorder = mid
		} else {
			r = mid - 1
		}
	}
	return leftBorder
}

func binarySearchRight(nums []int, target int) int {
	l, r := 0, len(nums)-1
	rightBorder := -1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			rightBorder = mid
			l = mid + 1
		}
	}
	return rightBorder
}
