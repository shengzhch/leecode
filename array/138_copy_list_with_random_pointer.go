package main

import (
	"fmt"
	"github.com/shengzhch/leecode/array/common"
)

func main() {
	a := common.RandomListNode{
		Val: 1,
	}
	b := common.RandomListNode{
		Val: 2,
	}
	c := common.RandomListNode{
		Val: 3,
	}

	a.Next = &b
	b.Next = &c
	a.Random = &c

	newl := DeepCopyWithRandomPoint(&a)
	fmt.Println(newl.Val)
	fmt.Println(newl.Next.Val)
	fmt.Println(newl.Next.Next.Val)
	fmt.Println(newl.Random.Val)
	a.Random = &b
	fmt.Println(newl.Random.Val)
}

//深度拷贝
func DeepCopyWithRandomPoint(cur *common.RandomListNode) *common.RandomListNode {
	var mapOfoldNode = make(map[*common.RandomListNode]int)
	var arrayOfnewNode = make([]*common.RandomListNode,4)
	var index = 0
	temp := cur

	//第一次遍历
	for temp != nil {
		var newNode = &common.RandomListNode{Val: temp.Val}
		mapOfoldNode[temp] = index
		arrayOfnewNode[index] = newNode
		temp = temp.Next
		index++
	}

	arrayOfnewNode[index] = &common.RandomListNode{}

	//第二次遍历
	temp = cur
	index = 0
	for temp != nil {
		arrayOfnewNode[index].Next = arrayOfnewNode[index+1]
		if temp.Random != nil {
			arrayOfnewNode[index].Random = arrayOfnewNode[mapOfoldNode[temp.Random]]
		}
		temp = temp.Next
		index++
	}
	return arrayOfnewNode[0]
}
