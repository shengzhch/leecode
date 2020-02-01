package common

import (
	"fmt"
)

func Min(vs ...int) int {
	if len(vs) == 0 {
		panic("func min args empty")
	}
	m := vs[0]
	for _, n := range vs[1:] {
		if n < m {
			m = n
		}
	}
	return m
}

func Max(vs ...int) int {
	if len(vs) == 0 {
		panic("func min args empty")
	}
	m := vs[0]
	for _, n := range vs[1:] {
		if n > m {
			m = n
		}
	}
	return m
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Reverse(nums []int) {
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}

// 递归版的二分查找虽高效，但调用传参时注意，nums 为空时 l,r 应为 0,-1，否则将 panic
func BinarySearch(nums []int, l, r int, target int) int {
	if l > r {
		return -1
	}

	mid := (l + r) / 2
	switch {
	case target < nums[mid]:
		return BinarySearch(nums, l, mid-1, target)
	case target > nums[mid]:
		return BinarySearch(nums, mid+1, r, target)
	default:
		return mid
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type RandomListNode struct {
	Val  int
	Next *RandomListNode
	//随机指向一个结点
	Random *RandomListNode
}

func NewList(nums []int) *ListNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0], Next: nil}
	cur := head
	for i := 1; i < n; i++ {
		newNode := &ListNode{Val: nums[i], Next: nil}
		cur.Next = newNode
		cur = newNode
	}
	return head
}

func (cur *ListNode) String() string {
	counts := 0
	var nums string
	var temp = new(ListNode)
	temp = cur
	for temp != nil {
		if counts == 0 {
			nums = nums + fmt.Sprintf("%v", temp.Val)
		} else {
			nums = nums + fmt.Sprintf(" -> %v", temp.Val)
		}
		counts++
		temp = temp.Next
	}
	return fmt.Sprintf("%d nodes: %v", counts, nums)
}

func (cur *ListNode) StringHasRing() string {
	counts := 0
	var nums string
	var temp = new(ListNode)
	var HasPrint = make(map[*ListNode]struct{})
	temp = cur
	for temp != nil {
		if counts == 0 {
			nums = nums + fmt.Sprintf("%v", temp.Val)
		} else {
			nums = nums + fmt.Sprintf(" -> %v", temp.Val)
		}
		if _, ok := HasPrint[temp]; ok {
			break
		}
		HasPrint[temp] = struct{}{}
		counts++
		temp = temp.Next
	}
	return fmt.Sprintf("%d nodes: %v", counts, nums)
}

func (cur *ListNode) Reverse() *ListNode {
	var newHead *ListNode = nil
	for cur != nil {
		next := cur.Next
		cur.Next = newHead
		newHead = cur
		cur = next
	}
	return newHead
}

func (cur *ListNode) ReverseM2N(m int, n int) *ListNode {
	changeLen := n - m + 1
	var preChangeHead, changeHead, changeTail *ListNode

	var newHead = cur

	for cur != nil && m > 1 {
		preChangeHead = cur
		cur = cur.Next
		m--
	}

	changeTail = cur
	for cur != nil && changeLen > 0 {
		next := cur.Next
		cur.Next = changeHead
		changeHead = cur
		cur = next
		changeLen--
	}
	changeTail.Next = cur

	if preChangeHead != nil {
		preChangeHead.Next = changeHead
	} else {
		newHead = changeHead
	}
	return newHead
}
