package main

import (
	"fmt"
	"github.com/shengzhch/leecode/array/common"
)

func main() {
	l := common.NewList([]int{1, 2, 3, 4, 5, 6, 7})
	l.Next.Next.Next.Next.Next.Next.Next = l.Next.Next
	fmt.Println(l.StringHasRing())

	has, cur := FirstNodeOfListHasRing(l)

	fmt.Println("has:", has, ";ring: ", cur.StringHasRing())
}

//快慢指针(快指针走的距离是慢指针走的两倍) 环的开始结点：分别从相遇结点，头节点出发，相遇的结点即为开始的结点

//设开始结点到环的开始结点距离为a，环开始结点到相遇结点距离为b，相遇结点到环开始结点距离为c
//则a+b+c+b = 2（a+b） 结论 ==》 a = c
func FirstNodeOfListHasRing(cur *common.ListNode) (bool, *common.ListNode) {
	fast := cur
	slow := cur

	var meet *common.ListNode

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		} else {
			return false, nil
		}
		if fast == slow {
			meet = fast
			break
		}
	}

	for {
		if cur == meet {
			return true, cur
		}
		cur = cur.Next
		meet = meet.Next
	}
}
