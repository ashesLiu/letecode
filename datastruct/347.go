package datastruct

import (
	"container/heap"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	return topKHeap(nums, k)
}

var buckets map[int]int

type MyHeap struct{
	sort.IntSlice
}

func (hp *MyHeap)Less(i ,j int)bool{
	return buckets[hp.IntSlice[i]] < buckets[hp.IntSlice[j]]
}

func (hp *MyHeap)Push(x interface{}){
	hp.IntSlice = append(hp.IntSlice, x.(int))
}

func (hp *MyHeap)Pop() interface{}{
	n := len(hp.IntSlice)
	x := hp.IntSlice[n-1]
	hp.IntSlice = hp.IntSlice[:n-1]
	return x
}

func topKHeap(nums []int, k int) []int {
	buckets = make(map[int]int, k)
	for i := 0; i < len(nums); i++ {
		buckets[nums[i]]++
	}
	kHeap := &MyHeap{
		make([]int, 0, k),
	}
	heap.Init(kHeap)
	for key := range buckets {
		if len(kHeap.IntSlice) < k{
			heap.Push(kHeap, key)
		}else if buckets[key]>buckets[kHeap.IntSlice[0]]{
			heap.Pop(kHeap)
			heap.Push(kHeap, key)
		}
	}
	return kHeap.IntSlice
}
