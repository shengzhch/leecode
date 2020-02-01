package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{1, 2, 3, 4}))
}

func subsets(nums []int) [][]int {
	n := len(nums)
	seqs := make([][]int, 1)
	seqs = append(seqs, []int{nums[0]})

	for i := 1; i < n; i++ {
		for _, seq := range seqs {
			nextSeq := append(seq,nums[i])
			seqs = append(seqs, nextSeq)
		}
	}

	return seqs
}
