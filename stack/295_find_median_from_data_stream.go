package main

import (
	"fmt"
	"github.com/shengzhch/book_sfjj/common"
)

//支持添加,支持返回当前的中位数（排序后,奇数个数返回中间值,偶数个数返回中间两个的平均值）
//维护一个最大堆和一个最小堆，且最大堆和最小堆的个数相差不超过1，且最大堆的堆顶比最小堆的堆顶小
type MyMedianArr struct {
	minHeap *common.Heap
	maxHeap *common.Heap
	compare func(key1, key2 interface{}) int
}

func NewMyMedianArr(compare func(key1, key2 interface{}) int) *MyMedianArr {
	f1 := func(key1, key2 interface{}) int {
		return compare(key1, key2)
	}

	f2 := func(key1, key2 interface{}) int {
		return 0 - compare(key1, key2)
	}

	minH := &common.Heap{}
	maxH := &common.Heap{}
	minH.Init(f2)
	maxH.Init(f1)

	return &MyMedianArr{
		minHeap: minH,
		maxHeap: maxH,
		compare: compare,
	}
}

func (m *MyMedianArr) Add(n int) {
	if m.maxHeap.Size() == 0 {
		m.maxHeap.Insert(n)
		return
	}

	if m.minHeap.Size() == m.maxHeap.Size() {
		if m.compare(n, m.maxHeap.Top()) > 0 {
			m.minHeap.Insert(n)
		} else {
			m.maxHeap.Insert(n)
		}
	} else if m.minHeap.Size() < m.maxHeap.Size() {
		//minHeap中个数少,和最大堆堆顶元素比，minHeap插入两个中大的元素
		if m.compare(n, m.maxHeap.Top()) > 0 {
			m.minHeap.Insert(n)
		} else {
			m.minHeap.Insert(m.maxHeap.Extract())
			m.maxHeap.Insert(n)
		}
	} else {
		//maxHeap中个数少,和最小堆堆顶元素比，maxHeap插入两个中小的元素
		if m.compare(n, m.minHeap.Top()) > 0 {
			m.maxHeap.Insert(m.minHeap.Extract())
			m.minHeap.Insert(n)
		} else {
			m.maxHeap.Insert(n)
		}
	}
}

func (m *MyMedianArr) GetMedian() float64 {
	if m.minHeap.Size() > m.maxHeap.Size() {
		return float64(m.minHeap.Top().(int))
	}
	if m.minHeap.Size() < m.maxHeap.Size() {
		return float64(m.maxHeap.Top().(int))
	}

	return float64(m.minHeap.Top().(int)+m.maxHeap.Top().(int)) / 2.0

}

func main() {
	bigIsBig := func(key1, key2 interface{}) int {
		if key1.(int) > key2.(int) {
			return 1
		}
		if key1.(int) == key2.(int) {
			return 0
		}
		return -1
	}

	arr := NewMyMedianArr(bigIsBig)

	test := []int{10, 30, 50, 20, 40, 60}

	for _, v := range test {
		arr.Add(v)
		//fmt.Println("now max ", arr.maxHeap.Data())
		//fmt.Println("now min ", arr.minHeap.Data())
		fmt.Println("now median ", arr.GetMedian())

	}
}

//10 == 10
//10 30  ==> 20
//10 30 50  ==> 30
//10 20 30 50 ==> 25
//10 20 30 40 50 ==> 30
//10 20 30 40 50 60  == 35