package main

import "fmt"

/*

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

*/

func twosum(a []int, target int) []int {
	v2i := make(map[int]int, len(a))
	for i, v := range a {
		del := target - v
		if rel, ok := v2i[del]; ok {
			return []int{rel, i}
		}
		v2i[v] = i
	}
	return nil
}

func main() {
	fmt.Println("rel", twosum([]int{2, 7, 11, 15}, 9))
}
