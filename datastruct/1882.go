type server struct{
    weight int
    idx int
    inTime int
}


type serverHeap []*server

func (s serverHeap)Len()int{
    return len(s)
}
func (s serverHeap)Swap(i, j int){
    s[i], s[j] = s[j], s[i]
}
func (s serverHeap)Less(i, j int)bool{
    if s[i].inTime == s[j].inTime {
		if s[i].weight == s[j].weight {
			return s[i].idx < s[j].idx
		}
		return s[i].weight < s[j].weight
	}
	return s[i].inTime < s[j].inTime
}
func (s *serverHeap)Push(val interface{}){
    *s = append(*s, val.(server))
}
func (s *serverHeap)Pop() interface{}{
    old := *s
    n := len(old)
    ans := old[n-1]
    old[n-1] = nil
    *s = old[0:n-1]
    return ans
}

func assignTasks(servers []int, tasks []int) []int {
    n:= len(tasks)
    // queue := make([]int, n)
    startTime := make([]int, n)
    ans := make([]int, n)
    for i:=0;i<n;i++{
        startTime[i] = i
        ans[i] = -1
    }
    serverArray := make([]server, len(servers))
    for i:=0;i<len(servers);i++{
        serverArray[i] = server{weight: servers[i], idx: i}
    }
    sort.Slice(serverArray, func (i, j int) bool{
		a := serverArray[i]
		b := serverArray[j]
        if a.weight == b.weight{
            return a.idx < b.idx
        }else{
            return a.weight < b.weight
        }
    })
    for i:=0;i<n;{
        for j:=range serverArray{
            if serverArray[j].inTime <= startTime[i]{
                ans[i] = serverArray[j].idx
                serverArray[j].inTime = startTime[i] + tasks[i]
                break
            }
        }
        if ans[i] == -1{
            min := serverArray[0].inTime
            for j:=range serverArray{
                if serverArray[j].inTime < min{
                    min = serverArray[j].inTime
                }
            }
            startTime[i] = min
        }else{
            i++
        }
    }


    return ans
}