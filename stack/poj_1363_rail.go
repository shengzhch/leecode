package main

import (
	"fmt"
	"github.com/shengzhch/book_sfjj/common"
)

//判断出栈顺序是否合法，数字入栈后可以立即出栈也可以等待其他数字入栈出栈后进行出栈

func check_is_valid_order(q *common.Queue, n int) bool {
	s := &common.Stack{}
	for i := 1; i <= n; i++ {
		s.Push(i)
		//有序
		if !s.IsEmpty() && q.Head().(int) < s.Top().(int) {
			return false
		}
		for (!s.IsEmpty() && q.Head() == s.Top()) {
			s.Pop()
			q.DeQueue()
		}
	}

	if q.Size() > 0 {
		return false
	}
	return true
}

func main() {

	q := &common.Queue{}
	q.EnQueue(3)
	q.EnQueue(2)
	q.EnQueue(5)
	q.EnQueue(4)
	q.EnQueue(1)

	fmt.Println(check_is_valid_order(q,5))

	q = &common.Queue{}
	q.EnQueue(3)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(4)
	q.EnQueue(5)

	fmt.Println(check_is_valid_order(q,5))
}
