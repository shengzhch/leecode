package main

import (
	"fmt"
	"github.com/shengzhch/leecode/array/common"
)

func MaxNoDecSunstring(s string) (int, int) {
	var size, i, j, k, max int;
	size = len(s)
	var start int

	/* i 是外层循环，j是下一次开始循环的起点，当新来的i有重复的时，更新j，并判断是否需要更新最大值*/
	for i = 0; i < size; i++ {
		for k = j; k < i; k++ {
			if (s[k] == s[i]) {
				j = k + 1;
				break;
			}
		}
		if (i-j+1 > max) {
			start = j
			max = i - j + 1
		}
		fmt.Println(" j max", j, max)
	}
	return max, start;
}

func MaxNoDecSunstring2(s string) (int, int) {
	st := make(map[int32]int)
	var i, max, start int
	for k, v := range s {
		if rel, ok := st[v]; ok {
			i = common.Max(rel, i)
		}
		if k-i+1 > max {
			max = k - i + 1
			start = i
		}
		st[v] = k + 1

		fmt.Println(" i,k max", i,k, max)

	}
	return max, start;
}

func main() {
	str := "abcadcb"
	//m, s := MaxNoDecSunstring(str)
	//fmt.Println("m s", m, s, "  -- ", str[s:s+m])

	m, s := MaxNoDecSunstring2(str)
	fmt.Println("m s", m, s, )
}
