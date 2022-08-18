package array

import "fmt"

func totalFruit(fruits []int) int {
	bucket := make(map[int]int)
	ans := 0
	l, r := 0, 0
	fmt.Printf("%v\n", fruits)
	for r < len(fruits) {
		if _, ok := bucket[fruits[r]]; len(bucket) < 2 || ok {
			bucket[fruits[r]] = r
			r++
		} else {
			if ans < r-l {
				ans = r - l
			}
			for l != bucket[fruits[l]] {
				l++
			}
			delete(bucket, fruits[l])
			l++
			// fmt.Printf("shrink:%d %d\n", l, r)
		}
	}
	return IntMax(ans, r-l)
}

func IntMax(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
