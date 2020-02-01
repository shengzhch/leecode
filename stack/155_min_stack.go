package main

import (
	"fmt"
	"github.com/shengzhch/book_sfjj/common"
)

type MyMinstack struct {
	data *common.Stack
	min  *common.Stack
}

func (s *MyMinstack) Init() {
	s.data.Init()
	s.min.Init()
}

func (s *MyMinstack) Push(data interface{}, cf common.CF) {
	if s.min.Size() == 0 {
		s.min.Push(data)
	} else if cf(data, s.Getmin()) <= 0 {
		s.min.Push(data)
	} else {
		s.min.Push(s.Getmin())
	}

	s.data.Push(data)
}

func (s *MyMinstack) Pop() interface{} {
	pv := s.data.PopValue()
	s.min.PopValue()
	return pv
}

func (s *MyMinstack) Getmin() interface{} {
	return s.min.Top()
}

func (s *MyMinstack) Show() {
	fmt.Println("data --- ")
	tf := func(e *common.ListElm, args ...interface{}) bool {
		if e == nil {
			return true
		}
		fmt.Println(e.GetValue())
		return false
	}
	(*common.List)(s.data).Traverse(tf)

	fmt.Println("min --- ")
	(*common.List)(s.min).Traverse(tf)
}

func main() {
	cf := common.Compare_int
	ms := &MyMinstack{
		&common.Stack{}, &common.Stack{},
	}
	ms.Init()

	ms.Push(3, cf)
	ms.Push(2, cf)
	ms.Push(2, cf)
	ms.Push(1, cf)
	ms.Push(4, cf)
	ms.Push(3, cf)

	ms.Show()
	fmt.Println("ms.Pop() ", ms.Pop())
	fmt.Println("ms.Pop() ", ms.Pop())
	fmt.Println("ms.Getmin() ", ms.Getmin())
	fmt.Println("ms.Pop() ", ms.Pop())
	fmt.Println("ms.Getmin() ", ms.Getmin())
	fmt.Println("ms.Pop() ", ms.Pop())
	ms.Show()
}
