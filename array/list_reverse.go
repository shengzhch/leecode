package main

import (
	"fmt"
	"github.com/shengzhch/leecode/array/common"
)

func main() {
	l := common.NewList([]int{1, 2, 3, 4, 5})
	fmt.Println("before reverse ", l.String())
	newl := l.Reverse()
	fmt.Println("after reverse ", newl.String())
	l = common.NewList([]int{1, 2, 3, 4, 5})
	newl2 := l.ReverseM2N(2, 4)
	fmt.Println("after reverse m-n", newl2.String())
}
