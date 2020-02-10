package main

import (
	"fmt"
	"github.com/shengzhch/book_sfjj/common"
	"strconv"
)

//字符串表示一个数字num，将num中k的数字移除，求移除k个数字后，可以获得最小可能的新数字。数字不为0开头
//num "1432219", k = 3 ==> 1219 最小

//思路 去掉一个数字，尽可能的让最高位最小，最高次位最小

func removeKdigits(arr []int, k int) []int {
	l := len(arr)
	s := common.Stack{}
	s.Init()
	for i := 0; i < l; i++ {
		number := arr[i]
		for s.Size() > 0 && number < s.Top().(int) && k > 0 {
			//出栈
			s.Pop()
			k--
		}

		//if number != 0 || s.Size() != 0 { //栈为空且number为0，表示最高位为0不起作用，
		//	s.Push(number)
		//}

		s.Push(number)
	}

	for s.Size() != 0 && k > 0 {
		s.Pop()
		k--
	}
	size := s.Size()
	list := common.List(s)
	rel := make([]int, 0, size)

	list.Traverse(func(e *common.ListElm, args ...interface{}) bool {
		*args[0].(*[]int) = (append(*args[0].(*[]int), e.GetValue().(int)))
		return false
	}, &rel)

	return rel
}

func main() {
	s := "1432219"
	arr := make([]int, 0, len(s))
	for _, v := range s {
		num, _ := strconv.Atoi(string(v))
		arr = append(arr, num)
	}
	rel := removeKdigits(arr, 3)
	fmt.Println("rel ", rel)

	after := 0
	for _, v := range rel {
		after = after*10 + v
	}
	fmt.Println("after ", after)
}
