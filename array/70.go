package main

import "fmt"

func main() {
	fmt.Println(climbStairs1(5)) // 3
	fmt.Println(climbStairs2(5)) // 5
	fmt.Println(climbStairs3(5)) // 8
}

func climbStairs1(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs1(n-1) + climbStairs1(n-2) // 存在大量重复计算
}

func climbStairs2(n int) int {
	steps := make([]int, n+1)
	steps[1], steps[2] = 1, 2 // 边界

	for i := 3; i <= n; i++ {
		steps[i] = steps[i-2] + steps[i-1] // 状态转移
	}

	return steps[n] // 最优子结构
}

func climbStairs3(n int) int {
	if n == 1 {
		return 1
	}

	i, j := 1, 2
	for n > 2 {
		i, j = j, i+j // bottom -> up
		n--
	}
	return j
}


