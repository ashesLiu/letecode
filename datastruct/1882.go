type server struct{
    weight int
    idx int
    inTime int
}


type baseHeap []*server
type workHeap struct{
    baseHeap
}
type freeHeap struct{
    baseHeap
}

func (b baseHeap) Len()int{
    return len(b)
}

func (b baseHeap) Swap(i,j int){
    b[i], b[j] = b[j],b[i]
}

func (w workHeap) Less(i, j int) bool{
    p, q:= w.baseHeap[i], w.baseHeap[j]
    if p.inTime == q.inTime{
        if p.weight == q.weight{
            return p.idx < q.idx
        }
        return p.weight < q.weight
    }
    return p.inTime < q.inTime
}

func (f freeHeap) Less(i,j int) bool{
    p, q:= f.baseHeap[i], f.baseHeap[j]
    if p.weight == q.weight{
        return p.idx < q.idx
    }
    return p.weight < q.weight
}

func (b *baseHeap) Push(val interface{}){
    *b = append(*b, val.(*server))
}

func (b *baseHeap) Pop() interface{}{
    old := *b
    n := len(old)
    ans := old[n-1]
    old[n-1] = nil
    *b = old[0:n-1]
    return ans
}

func max(a, b int)int{
    if a < b {
        return b
    }
    return a
}

func assignTasks(servers []int, tasks []int) []int {
    n:= len(tasks)
    m:= len(servers)
    wh := workHeap{make(baseHeap,0)}
    fh := freeHeap{make(baseHeap,m)}
    ans := make([]int, n)
    for j:= range servers{
        fh.baseHeap[j] = &server{servers[j],j,0}
    }
    heap.Init(&fh)
    heap.Init(&wh)
    ts := 0
    for i:=0;i<n;{
        ts = max(ts,i)
        for len(wh.baseHeap)>0 && wh.baseHeap[0].inTime<=ts {
            s := heap.Pop(&wh).(*server)
            heap.Push(&fh, s)
        }
        if len(fh.baseHeap)==0{
            ts=wh.baseHeap[0].inTime
        }else{
            s := heap.Pop(&fh).(*server)
            ans[i] = s.idx
            s.inTime = ts + tasks[i]
            heap.Push(&wh, s)
            i++
        }
    }
    return ans
}type server struct{
    weight int
    idx int
    inTime int
}


type baseHeap []*server
type workHeap struct{
    baseHeap
}
type freeHeap struct{
    baseHeap
}

func (b baseHeap) Len()int{
    return len(b)
}

func (b baseHeap) Swap(i,j int){
    b[i], b[j] = b[j],b[i]
}

func (w workHeap) Less(i, j int) bool{
    p, q:= w.baseHeap[i], w.baseHeap[j]
    if p.inTime == q.inTime{
        if p.weight == q.weight{
            return p.idx < q.idx
        }
        return p.weight < q.weight
    }
    return p.inTime < q.inTime
}

func (f freeHeap) Less(i,j int) bool{
    p, q:= f.baseHeap[i], f.baseHeap[j]
    if p.weight == q.weight{
        return p.idx < q.idx
    }
    return p.weight < q.weight
}

func (b *baseHeap) Push(val interface{}){
    *b = append(*b, val.(*server))
}

func (b *baseHeap) Pop() interface{}{
    old := *b
    n := len(old)
    ans := old[n-1]
    old[n-1] = nil
    *b = old[0:n-1]
    return ans
}

func max(a, b int)int{
    if a < b {
        return b
    }
    return a
}

func assignTasks(servers []int, tasks []int) []int {
    n:= len(tasks)
    m:= len(servers)
    wh := workHeap{make(baseHeap,0)}
    fh := freeHeap{make(baseHeap,m)}
    ans := make([]int, n)
    for j:= range servers{
        fh.baseHeap[j] = &server{servers[j],j,0}
    }
    heap.Init(&fh)
    heap.Init(&wh)
    ts := 0
    for i:=0;i<n;{
        ts = max(ts,i)
        for len(wh.baseHeap)>0 && wh.baseHeap[0].inTime<=ts {
            s := heap.Pop(&wh).(*server)
            heap.Push(&fh, s)
        }
        if len(fh.baseHeap)==0{
            ts=wh.baseHeap[0].inTime
        }else{
            s := heap.Pop(&fh).(*server)
            ans[i] = s.idx
            s.inTime = ts + tasks[i]
            heap.Push(&wh, s)
            i++
        }
    }
    return ans
}