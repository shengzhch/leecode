package main

import (
	"fmt"
	"github.com/shengzhch/leecode/array/common"
)

func main() {
	l := common.NewList([]int{1, 4, 3, 7, 2, 6, 5})

	newL := PartionList(l, 4)
	fmt.Println("after partion[4] ", newL.String())

}

func PartionList(cur *common.ListNode, n int) *common.ListNode {
	lessHead := common.ListNode{}
	moreHead := common.ListNode{}
	lessPtr := &lessHead
	morePtr := &moreHead

	for cur != nil {
		if cur.Val < n {
			lessPtr.Next = cur
			lessPtr = lessPtr.Next
		} else {
			morePtr.Next = cur
			morePtr = morePtr.Next
		}
		cur = cur.Next
	}

	lessPtr.Next = moreHead.Next
	return lessHead.Next
}
