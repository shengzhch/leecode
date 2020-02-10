package main

import (
	"fmt"
	"github.com/shengzhch/book_sfjj/common"
	"sort"
)

type Pair struct {
	x1 int //加油站到终点的距离
	x2 int //加油站的油量
}

type Pairs struct {
	bs []Pair
}

func (b *Pairs) Len() int {
	return len(b.bs)
}

func (b *Pairs) Less(i, j int) bool {
	return b.bs[i].x1 > b.bs[j].x1
}

func (b *Pairs) Swap(i, j int) {
	b.bs[i], b.bs[j] = b.bs[j], b.bs[i]
}

var bigIsBig = func(key1, key2 interface{}) int {
	if key1.(int) > key2.(int) {
		return 1
	}
	if key1.(int) == key2.(int) {
		return 0
	}
	return -1
}

//l 为起点到终点的距离，p为起点初始油量
func getMinimumStop(l int, p int, ps *Pairs) (int, []int) {
	var rel = 0
	trace := make([]int, 0)

	q := &common.Heap{}
	//最大堆
	q.Init(bigIsBig)

	//添加终点pair
	ps.bs = append(ps.bs, Pair{
		x1: 0,
		x2: 0,
	})

	//从大到小排序
	sort.Sort(ps)

	fmt.Println("ps.bs",ps.bs)

	n := len(ps.bs)

	for i := 0; i < n; i++ {
		//到该点要走的距离
		dis := l - ps.bs[i].x1

		//需要加油
		for q.Size() > 0 && p < dis {
			oil := q.Top().(int)
			p = p + oil
			trace = append(trace, oil)
			q.Extract()
			rel++
		}

		//无法到达
		if q.Size() == 0 && p < dis {
			return -1, nil
		}

		//油量更新
		p = p - dis
		//到终点距离更新
		l = ps.bs[i].x1
		//添加该点油量
		q.Insert(ps.bs[i].x2)
	}

	return rel, trace
}

func main() {
	l := 25
	p := 16
	//4个加油站
	ps := &Pairs{bs: []Pair{
		{
			x1: 4,
			x2: 4,
		},
		{
			x1: 10,
			x2: 3,
		},
		{
			x1: 11,
			x2: 5,
		},
		{
			x1: 15,
			x2: 2,
		},
	}}

	rel, stop := getMinimumStop(l, p, ps)

	fmt.Println("rel and stop", rel, stop)
}
