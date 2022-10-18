package datastruct

type UnionFind struct{
	father []int
	weight []float64
}

func NewUnionFind(size int) *UnionFind{
	father := make([]int, size)
	weight := make([]float64, size)
	for i:=range father{
		father[i]=i
		weight[i]=1
	}
    return &UnionFind{father, weight}
}

func (u *UnionFind) find(num int) int{
	if num == u.father[num]{
        return num
    }
    x := u.find(u.father[num])
    u.weight[num] *= u.weight[u.father[num]]
    u.father[num] = x
    return x
}

func (u *UnionFind) union(x int, y int, v float64){
	// x = fx*w[x]
	// y = fy*w[y]
	// x/y = v
	// fx*w[x] / (fy*w[y]) = v
	// fx/fy = v/w[x]*w[y]
	fx, fy := u.find(x), u.find(y)
	if fx != fy{
		u.father[fx] = fy
		u.weight[fx] = v/u.weight[x]*u.weight[y]
	}
}



func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	idmap := make(map[string]int)
	for i:= range equations{
		getId(idmap, equations[i][0])
		getId(idmap, equations[i][1])
	}
	group := NewUnionFind(len(idmap))
	for i, eq:=range equations{
		group.union(idmap[eq[0]], idmap[eq[1]], values[i])
	}
    fmt.Println(group.father)
    fmt.Println(group.weight)
	ans := make([]float64, len(queries))
	for i, qu:=range queries{
        idA, ok := idmap[qu[0]]
        idB, ok2 := idmap[qu[1]]
		if ok&&ok2&&group.find(idA) == group.find(idB){
			ans[i] = group.weight[idA] / group.weight[idB]
		}else{
			ans[i] = -1
		}
	}
	return ans
}



func getId(idmap map[string]int, key string){
	if _, ok:= idmap[key]; !ok{
		idmap[key] = len(idmap)
	}
}
