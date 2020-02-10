package main

import "fmt"

//tot 数组arr 从第i个位置最多可以跳arr[i]步，判断能否从第0个位置能否跳到最后位置。
// 从0位置可以最远可以跳到第i个位置，之间的都可以，从0-i中选择可以跳到更远位置的下标。

func canJump(arr []int) bool {
	l := len(arr)
	indexArr := make([]int, l)
	for i := 0; i < l; i++ {
		indexArr[i] = i + arr[i]
	}
	jump := 0
	maxIndex := indexArr[0]


	for jump < l && jump <= maxIndex {
		if maxIndex < indexArr[jump] {
			maxIndex = indexArr[jump]
		}
		jump++
	}
	//可以达到数组尾部
	if jump == l {
		return true
	}
	return false
}

//即每次选择可以跳到位置最大的下标
func minJump(arr []int) (int, []int) {
	l := len(arr)
	if l < 2 {
		return 0, nil
	}

	curMaxIndex := arr[0]    //当前可达到的最远位置
	preMaxMaxIndex := arr[0] //遍历各个位置中可达到的最远位置
	jumpMin := 1

	trace := make([]int, 0)
	var jumpIndex = 0
	trace = append(trace, jumpIndex)
	for i := 1; i < l; i++ {
		if i > curMaxIndex {
			trace = append(trace, jumpIndex)
			jumpMin++
			curMaxIndex = preMaxMaxIndex
		}
		if preMaxMaxIndex < arr[i]+i {
			preMaxMaxIndex = arr[i] + i
			jumpIndex = i
		}
	}

	return jumpMin, trace
}

func main() {
	arr := []int{2, 3, 1, 1, 4}
	fmt.Println("rel ", canJump(arr))

	min, trace := minJump(arr)

	fmt.Println("min and trace ", min, trace)
}
