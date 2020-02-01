package main

import (
	"fmt"
	"github.com/shengzhch/leecode/array/common"
)


/*
记录当前最大值，然后向中间移动。计算新的面积与旧的面积做比较
*/
func maxArea(height []int) int {
	maxMulti := 0
	left, right := 0, len(height)-1
	for left < right {
		w := right - left
		h := common.Min(height[left], height[right])
		maxMulti = common.Max(maxMulti, w*h)
		if height[left] <= height[right] {
			left++ // 往右边走找更长的线
		} else {
			right-- // 往左边走
		}
	}
	return maxMulti
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
