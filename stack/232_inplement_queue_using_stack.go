package main

import (
	"fmt"
	"github.com/shengzhch/book_sfjj/common"
)

//用栈实现队列 push顺序 1,2,3,4,5  ==> 队列 5，4，3，2，1  栈 1，2，3，4，5
type Myqueue struct {
	push *common.Stack
	pop  *common.Stack
}

func (s *Myqueue) EnQueue(val interface{}) {
	s.push.Push(val)
}

func (s *Myqueue) DeQueueWithValue() interface{} {
	if s.pop.Size() > 0 {
		return s.pop.PopValue()
	}
	for s.push.Size() > 0 {
		s.pop.Push(s.push.PopValue())
	}
	return s.pop.PopValue()
}

func (s *Myqueue) Head() interface{} {
	if s.pop.Size() > 0 {
		return s.pop.Top()
	}
	for s.push.Size() > 0 {
		s.pop.Push(s.push.PopValue())
	}
	return s.pop.Top()

}
func (s *Myqueue) Empty() bool {
	return s.push.Size()+s.pop.Size() <= 0
}

func main() {
	s := &Myqueue{push: &common.Stack{}, pop: &common.Stack{}}
	s.EnQueue(1)
	s.EnQueue(2)
	s.EnQueue(3)
	s.EnQueue(4)
	s.EnQueue(5)

	fmt.Println("s.DeQueueWithValue ", s.DeQueueWithValue())
	fmt.Println("s.DeQueueWithValue ", s.DeQueueWithValue())
	fmt.Println("s.Head", s.Head())
	fmt.Println("s.Empty", s.Empty())
	fmt.Println("s.DeQueueWithValue ", s.DeQueueWithValue())
	fmt.Println("s.DeQueueWithValue ", s.DeQueueWithValue())
	fmt.Println("s.DeQueueWithValue ", s.DeQueueWithValue())
	fmt.Println("s.Empty", s.Empty())
}
