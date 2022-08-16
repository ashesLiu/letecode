package array

import (
	"fmt"
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
	arr := []int{5, 7, 3, 9, 5, 8}
	QuichSort(arr, 0, len(arr))
	fmt.Printf("arr: %v\n", arr)
}
