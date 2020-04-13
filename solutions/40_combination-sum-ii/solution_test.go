package _0_combination_sum_ii

import (
	"fmt"
	"testing"
)

func combinationSum3(candidates []int, target int) [][]int {
	quickSort(candidates)
	var res [][]int
	for i, c := range candidates {
		if i > 0 && c == candidates[i-1] {
			continue
		}
		if c > target {
			return res
		}
		if c == target {
			res = append(res, []int{c})
			return res
		}
		for _, sol := range combinationSum2(candidates[i+1:], target-c) {
			res = append(res, append([]int{c}, sol...))
		}
	}
	return res
}

func combinationSum2(candidates []int, target int) [][]int {
	quickSort(candidates)
	return solution(candidates, target)
}

func solution(candidates []int, target int) [][]int {
	var res [][]int
	for i, c := range candidates {
		if i > 0 && c == candidates[i-1] {
			continue
		}
		if c > target {
			return res
		}
		if c == target {
			res = append(res, []int{c})
			return res
		}
		for _, sol := range solution(candidates[i+1:], target-c) {
			res = append(res, append([]int{c}, sol...))
		}
	}
	return res
}

func quickSort(nums []int) {
	l := len(nums)
	if l < 2 {
		return
	}

	head, tail := 1, l-1
	privot := nums[0]
	for head < tail {
		for head < tail && nums[head] < privot {
			head++
		}
		for tail > head && nums[tail] >= privot {
			tail--
		}
		nums[head], nums[tail] = nums[tail], nums[head]
	}
	if nums[tail] < privot {
		nums[0], nums[tail] = nums[tail], nums[0]
	}

	quickSort(nums[:head])
	quickSort(nums[tail:])
}

func Test_(t *testing.T) {
	candidates := []int{10,1,2,7,6,1,5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}
