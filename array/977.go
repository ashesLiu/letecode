package array

<<<<<<< HEAD
=======
import (
	"math/rand"
)

>>>>>>> partitionv1.1
func SortedSquares(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		nums[i] = nums[i] * nums[i]
	}

	// 快排
	QuichSort(nums, 0, len(nums))
	return nums
}

// quichSort: [left,right)
func QuichSort(nums []int, left, right int) {
	if left < right {
		idx := partition(nums, left, right)
		QuichSort(nums, left, idx)
		QuichSort(nums, idx+1, right)
	}
}

func partition(nums []int, l, r int) int {
<<<<<<< HEAD
	// 选择nums[l]为枢轴元素
=======
	// 选择任意值为枢轴元素
	pivot := rand.Intn(r-l) + l
	swap(nums, pivot, l)
>>>>>>> partitionv1.1
	v := nums[l]
	// 双指针: i遍历 pos为已遍历元素的划分点
	// 循环不变量： nums[l:pos)<pivot nums[pos:i)>=pivot
	pos := l
	for i := l + 1; i < r; i++ {
		if nums[i] < v {
<<<<<<< HEAD
			swap(nums, pos, i)
			pos++
		}
	}
	// swap(nums, pos, l)
=======
			pos += 1
			swap(nums, pos, i)
		}
	}
	swap(nums, pos, l)
	return pos
}

func partitionv0(nums []int, l, r int) int {
	v := nums[l]
	// 双指针: i遍历 pos为已遍历元素的划分点
	// 循环不变量： nums[l:pos)<pivot nums[pos:i)>=pivot
	pos := l
	for i := pos; i < r; i++ {
		if nums[i] < v {
			pos += 1
			swap(nums, pos, i)
		}
	}
	swap(nums, pos, l)
>>>>>>> partitionv1.1
	return pos
}

func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}
