package datastruct

import (
	"fmt"
	"strings"
)

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	data := make(map[string]float64)
	for i := 0; i < len(equations); i++ {
		eq1, eq2 := code(equations[i][0], equations[i][1]), code(equations[i][1], equations[i][0])
		data[eq1] = values[i]
		data[eq2] = 1.0 / values[i]
	}

	var dfs func(known map[string]float64, left string, target string, value float64, visited map[string]bool) float64

	dfs = func(known map[string]float64, left string, target string, value float64, visited map[string]bool) float64 {
		if ans, ok := known[code(left, target)]; ok {
			return value * ans
		}
		res := -1.0
		for k, v := range known {
			eqs := decode(k)
			if eqs[0] == left && !visited[eqs[0]] {
				visited[eqs[0]] = true
				res = dfs(known, target, eqs[1], value*v, visited)
				if res != -1.0 {
					return res
				}
				visited[eqs[0]] = false
			}
		}
		return res
	}
	arr := make([]float64, len(queries))
	for i := range queries {
		visited := make(map[string]bool)
		arr[i] = dfs(data, queries[i][0], queries[i][1], 1, visited)
	}
	return arr
}

func code(a, b string) string {
	return fmt.Sprintf("%s/%s", a, b)
}

func decode(s string) []string {
	return strings.Split(s, "/")
}
