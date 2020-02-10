package main

import "fmt"

//tot 摇摆子序列最长长度(子序列非连续)
//

const (
	begin = 1
	up    = 2
	down  = 3
)

func wiggleMaxLength(arr []int) int {
	if len(arr) < 2 {
		return len(arr)
	}

	var state = begin
	var maxLength = 1

	//状态机
	for i := 1; i < len(arr); i++ {
		switch state {
		case begin:
			if arr[i-1] < arr[i] {
				state = up
				maxLength++
			} else if arr[i-1] > arr[i] {
				state = down
				maxLength++
			}
		case up:
			if arr[i-1] > arr[i] {
				state = down
				maxLength++
			}
		case down:
			if arr[i-1] < arr[i] {
				state = up
				maxLength++
			}
		}
	}
	return maxLength
}

func main() {
	arr := []int{1, 17, 5, 10, 12, 15, 10, 5, 16, 8}

	fmt.Println(wiggleMaxLength(arr))
}
