package datastruct

import (
	"container/heap"
	"sort"
)

type Heap struct {
	sort.IntSlice
}

var data []int

func (hp *Heap) Less(i, j int) bool {
	return data[hp.IntSlice[i]] < data[hp.IntSlice[j]]
}

func (hp *Heap) Push(x interface{}) {
	hp.IntSlice = append(hp.IntSlice, x.(int))
}

func (hp *Heap) Pop() interface{} {
	old := hp.IntSlice
	n := len(old)
	x := old[n-1]
	hp.IntSlice = hp.IntSlice[:n-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	data = nums
	kHeap := &Heap{make([]int, k)}
	ans := make([]int, len(nums)-k+1)
	for i := 0; i < k; i++ {
		kHeap.IntSlice[i] = i
	}
	heap.Init(kHeap)
	ans[0] = kHeap.IntSlice[0]
	for i := k; i < len(nums); i++ {
		heap.Push(kHeap, i)
		for kHeap.IntSlice[0] < i-k+1 {
			heap.Pop(kHeap)
		}
		ans[i-k+1] = kHeap.IntSlice[0]
	}
	return ans
}

// func maxSlidingWindow(nums []int, k int) []int {

// 	// 单调队列
// 	// 单调队列保存有希望成为最大值的元素。
// 	pos := 0
// 	ans := make([]int, 0, len(nums)-k+1)
// 	queue := make([]int, 0, k)
// 	for pos < len(nums) {

// 		if len(queue) > 0 && queue[0] <= pos-k {
// 			queue = queue[1:]
// 		}

// 		for len(queue) > 0 && nums[queue[len(queue)-1]] <= nums[pos] {
// 			queue = queue[:len(queue)-1]
// 		}
// 		queue = append(queue, pos)

// 		if pos >= k-1 {
// 			ans = append(ans, nums[queue[0]])
// 		}

// 		pos++
// 	}
// 	return ans
// }
