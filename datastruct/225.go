package datastruct

type MyStack struct {
    inQueue []int
    outQueue []int
}


func Constructor() MyStack {
    // inQueue = make([]int, 0)
    // outQueue = make([]int, 0)
    return MyStack{}
}


func (this *MyStack) Push(x int)  {
    if len(this.outQueue) == 1{
        this.inQueue = append(this.inQueue, this.outQueue[0])
        this.outQueue = nil
    }
    this.outQueue = append(this.outQueue, x)
}

func (this *MyStack)transfer(){
    if len(this.outQueue) == 0{
        last := len(this.inQueue)-1
        for i:=0;i<last;i++{
            this.inQueue = append(this.inQueue, this.inQueue[i])
        }
        this.outQueue = append(this.outQueue, this.inQueue[last])
        this.inQueue = this.inQueue[last+1:]
    }
}


func (this *MyStack) Pop() int {
    this.transfer()
    x := this.outQueue[0]
    this.outQueue = nil
    return x
}


func (this *MyStack) Top() int {
    this.transfer()
    return this.outQueue[0]
}


func (this *MyStack) Empty() bool {
    if len(this.inQueue)==0 && len(this.outQueue)==0{
        return true
    }else{
        return false
    }
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */