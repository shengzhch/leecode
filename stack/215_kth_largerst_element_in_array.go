package main

import (
	"fmt"
	"github.com/shengzhch/book_sfjj/common"
)

//维护K大小的最小堆，堆顶元素即为第K大的数
func findKthLargest(arr []int, kth int) int {
	heap := &common.Heap{}
	lessIsBig := func(key1, key2 interface{}) int {
		if key2.(int) > key1.(int) {
			return 1
		}
		if key2.(int) == key1.(int) {
			return 0
		}
		return -1
	}
	heap.Init(lessIsBig)

	for _, v := range arr {
		if heap.Size() < kth {
			heap.Insert(v)
			//fmt.Printf("insert[%v] ,now size: %v\n", v, heap.Size())
		} else {
			//fmt.Println("size: ", heap.Size(),heap.Data())
			if heap.Top().(int) < v {
				_ = heap.Extract()
				heap.Insert(v)
				//fmt.Printf("replace[%v] to [%v] ,now size: %v\n", extract, v, heap.Size())
			}
		}
	}

	return heap.Top().(int)
}

func main() {
	arr := []int{3, 2, 5, 1, 4}
	fmt.Println("1th ", findKthLargest(arr, 1))
	fmt.Println("2th ", findKthLargest(arr, 2))
	fmt.Println("3th ", findKthLargest(arr, 3))
	fmt.Println("4th ", findKthLargest(arr, 4))
	fmt.Println("5th ", findKthLargest(arr, 5))
}
