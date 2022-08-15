package array

import (
	"fmt"
	"reflect"
	"testing"
)

func Array() {
	var arr [][]int
	arr = append(arr, make([]int, 3))
	arr = append(arr, make([]int, 3))
	for i := range arr {
		subArr := arr[i]
		for j, v := range subArr {
			subArr[j] = j
			v = j
			fmt.Printf("[%d]:%v=%d,%v=%d\t", j, &subArr[j], subArr[j], &v, v)

		}
		fmt.Println()
	}

}

func AssertEqual(t *testing.T, got interface{}, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want '%v'", got, want)
	}
}
