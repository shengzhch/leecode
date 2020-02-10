package main

import (
	"fmt"
	"github.com/shengzhch/book_sfjj/common"
)

//tot 糖果和孩子排序

func toSliceOfInterface(args ...int) []interface{} {
	rel := make([]interface{}, 0, len(args))
	for _, v := range args {
		rel = append(rel, v)
	}
	return rel
}

//g 需求因子，s 糖果大小
func findContentChildren(g, s []int) int {
	gg := toSliceOfInterface(g...)
	ss := toSliceOfInterface(s...)

	common.Qksort2(gg, 0, len(gg)-1, common.CF_INT)
	common.Qksort2(ss, 0, len(ss)-1, common.CF_INT)

	fmt.Println("gg: ", gg, "\n", "ss: ", ss)

	var child = 0
	var cookie = 0

	for child < len(gg) && cookie < len(ss) {
		if common.CF_INT(gg[child], ss[cookie]) <= 0 {
			child++
		}
		cookie++
	}
	return child
}

func main() {
	g := []int{5, 10, 2, 9, 15, 9}
	s := []int{6, 1, 20, 3, 8}
	fmt.Println("rel ", findContentChildren(g, s))
}
