package main

import (
	"fmt"
	"github.com/shengzhch/leecode/array/common"
)

/*

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5

*/

func fincMiddleNum(a, b []int) float64 {
	return 0
}

//按顺序排序，找到第k个数
func findKNum(a, b []int, k int) int {
	n1 := len(a)
	n2 := len(b)

	if n1 > n2 {
		n1, n2 = n2, n1
		a, b = b, a
	}

	if n1 == 0 {
		return b[k-1]
	}

	if k == 1 {
		return common.Min(a[0], b[0])
	}

	k1 := common.Min(k/2, n1)
	k2 := k - k1

	/* 肯定确定的是k2>k1,表示分别从a，b截取的部分*/

	//剔除不符合要求的部分
	switch {
	case a[k1-1] < b[k2-1]:
		return findKNum(a[k1:], b, k2)
	case a[k1-1] > b[k2-1]:
		return findKNum(a, b[:k2-1], k1)
	default:
		return a[k1-1]
	}
}

func main() {
	rel := findKNum([]int{1, 3, 5}, []int{2, 4, 6}, 2)

	fmt.Println("rel ", rel)
}
