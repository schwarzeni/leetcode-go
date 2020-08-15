package mm0305

// 2020.08.15
// https://leetcode-cn.com/problems/sort-of-stacks-lcci/
type Stack struct {
    data []int
}

func (s *Stack) Push(val int) {
    s.data = append(s.data, val)
}

func (s *Stack) Peek() int {
    if s.IsEmpty() {
        panic("stack is empty")
    }
    return s.data[len(s.data) - 1]
}

func (s *Stack) Pop() int {
    if s.IsEmpty() {
        panic("stack is empty")
    }
    v := s.Peek()
    s.data = s.data[:len(s.data) - 1]
    return v
}

func (s *Stack) IsEmpty() bool{
    return len(s.data) == 0
}


type SortedStack struct {
    stack Stack
    helper Stack
}


func Constructor() SortedStack {
    return SortedStack{
        stack: Stack{},
        helper: Stack{},
    }
}


func (this *SortedStack) Push(val int)  {
    if this.IsEmpty() {
        this.stack.Push(val)
        return
    }
    for !this.stack.IsEmpty() && this.Peek() < val {
        this.helper.Push(this.stack.Pop())
    }
    this.stack.Push(val)
    for !this.helper.IsEmpty() {
        this.stack.Push(this.helper.Pop())
    }
}


func (this *SortedStack) Pop()  {
    if this.stack.IsEmpty() {
        return
    }
    this.stack.Pop()
}


func (this *SortedStack) Peek() int {
    if this.stack.IsEmpty() {
        return -1
    }
    return this.stack.Peek()
}


func (this *SortedStack) IsEmpty() bool {
    return this.stack.IsEmpty()
}


/**
 * Your SortedStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.IsEmpty();
 */
