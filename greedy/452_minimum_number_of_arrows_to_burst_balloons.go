package main

import (
	"fmt"
	"sort"
)

//tot 击穿一个气球的同时，尽可能的击穿更多的其他气球

type Ballon struct {
	x1 int
	x2 int
}

type Ballons struct {
	bs []Ballon
}

func (b *Ballons) Len() int {
	return len(b.bs)
}

func (b *Ballons) Less(i, j int) bool {
	return b.bs[i].x1 < b.bs[j].x1
}

func (b *Ballons) Swap(i, j int) {
	b.bs[i], b.bs[j] = b.bs[j], b.bs[i]
}

func funcMinArrowShots(b *Ballons) int {
	if b.Len() == 0 {
		return 0
	}

	//排序后的
	sort.Sort(b)

	shotNum := 1

	//要维护的射击区间
	//shotBegin := b.bs[0].x1
	shotEnd := b.bs[0].x2

	for i := 1; i < b.Len(); i++ {
		//该区间可以射击该气球
		if b.bs[i].x1 <= shotEnd {
			//shotBegin = b.bs[i].x1
			if shotEnd > b.bs[i].x2 {
				shotEnd = b.bs[i].x2
			}
		} else {
			shotNum++
			//shotBegin = b.bs[i].x1
			shotEnd = b.bs[i].x2
		}
	}

	return shotNum
}

func main() {
	b := &Ballons{bs: []Ballon{
		{10, 16}, {2, 8}, {1, 6}, {7, 12},
	}}

	fmt.Println("rel ", funcMinArrowShots(b))
}
