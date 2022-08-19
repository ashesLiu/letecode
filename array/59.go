package array

import "fmt"

func generateMatrix(n int) [][]int {
	if n == 0 {
		return nil
	}
	direction := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	pos, i, j := 0, 0, 0
	nextPos := func(i int) int {
		if i == 3 {
			return 0
		} else {
			return i + 1
		}
	}
	arr := make([][]int, n)
	vis := make([][]int, n)
	for k := 0; k < n; k++ {
		vis[k] = make([]int, n)
		arr[k] = make([]int, n)
	}
	// fmt.Println("%v",arr)
	for k := 1; k <= n*n; k++ {
		fmt.Printf("%v %v\n", i, j)
		arr[i][j] = k
		vis[i][j] = 1
		tempi := i + direction[pos][0]
		tempj := j + direction[pos][1]
		if tempi >= n || tempi < 0 || tempj >= n || tempj < 0 || vis[tempi][tempj] == 1 {
			pos = nextPos(pos)
		}
		i += direction[pos][0]
		j += direction[pos][1]
	}
	return arr
}
