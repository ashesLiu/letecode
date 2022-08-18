package array

import (
	"math/rand"
	"sort"
	"testing"
)

func TestBinarySearch704(t *testing.T) {
	length := 100
	arr := make([]int, length)
	for i := range arr {
		arr[i] = rand.Intn(length * 10)
	}
	sort.Ints(arr)
	want := rand.Intn(length)
	target := arr[want]
	index := search(arr, target)
	AssertEqual(t, index, want)
}

func BenchmarkBinarySearch704(b *testing.B) {
	length := 100
	arr := make([]int, length)
	for i := range arr {
		arr[i] = rand.Intn(length * 10)
	}
	sort.Ints(arr)
	want := rand.Intn(length)
	target := arr[want]
	for i := 0; i < b.N; i++ {
		search(arr, target)
	}
}

func TestSortedSequence(t *testing.T) {
	arr := []int{5, 5, 6, 1, 6}
	QuickSort(arr, 0, len(arr))
	t.Logf("%v", arr)
}

func TestMinSubArrayLen(t *testing.T) {
	arr := []int{2, 3, 1, 2, 4, 3}
	ans := minSubArrayLen(7, arr)
	t.Logf("%v", ans)
}

func TestSearchRange(t *testing.T) {
	arr := []int{5, 7, 7, 8, 8, 10}
	t.Logf("%v", searchRange(arr, 8))
}
