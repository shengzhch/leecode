package main

import (
	"fmt"
	"github.com/shengzhch/book_sfjj/common"
)

//用队列实现栈 push顺序 1,2,3,4,5  ==> 队列 5，4，3，2，1  栈 1，2，3，4，5
type Mystack struct {
	temp *common.Queue
	data *common.Queue
}

func (s *Mystack) Push(val interface{}) {
	s.temp.EnQueue(val)
	for s.data.Size() > 0 {
		s.temp.EnQueue(s.data.DeQueueWithValue())
	}

	s.temp, s.data = s.data, s.temp
}

func (s *Mystack) Pop() interface{} {
	return s.data.DeQueueWithValue()
}

func (s *Mystack) Top() interface{} {
	return s.data.Head()

}
func (s *Mystack) Empty() bool {
	return s.data.Size() <= 0
}

func main() {
	s := &Mystack{temp: &common.Queue{}, data: &common.Queue{}}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	fmt.Println("s.POP", s.Pop())
	fmt.Println("s.POP", s.Pop())
	fmt.Println("s.TOP", s.Top())
	fmt.Println("s.Empty", s.Empty())
	fmt.Println("s.POP", s.Pop())
	fmt.Println("s.POP", s.Pop())
	fmt.Println("s.POP", s.Pop())
	fmt.Println("s.Empty", s.Empty())
}
