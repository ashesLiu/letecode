package array

func generateMatrix(n int) [][]int {
	if n == 0 {
		return nil
	}
	direction := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	// i:0,1 j:2,3
	border := []int{0, n - 1, 0, n - 1}
	pos, i, j := 0, 0, 0
	nextPos := func(i int) int {
		if i == 3 {
			return 0
		} else {
			return i + 1
		}
	}
	arr := make([][]int, n)
	for k := 0; k < n; k++ {
		arr[k] = make([]int, n)
	}
	for k := 1; k <= n*n; k++ {
		// fmt.Printf("%v %v\n", i, j)
		arr[i][j] = k
		tempi := i + direction[pos][0]
		tempj := j + direction[pos][1]
		if tempi < border[0] || tempi > border[1] || tempj < border[2] || tempj > border[3] {
			switch pos {
			case 0:
				border[0]++
			case 1:
				border[3]--
			case 2:
				border[1]--
			case 3:
				border[2]++
			}
			pos = nextPos(pos)
		}
		i += direction[pos][0]
		j += direction[pos][1]
	}
	return arr
}
