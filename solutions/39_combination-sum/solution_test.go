package _9_combination_sum

import (
	"fmt"
	"testing"
)

func combinationSum(candidates []int, target int) [][]int {
	quickSort(candidates)
	return solution(candidates, target)
}

func solution(candidates []int, target int) [][]int {
	var res [][]int
	for i, c := range candidates {
		if c > target {
			return res
		}
		if c == target {
			res = append(res, []int{c})
			return res
		}
		for _, sol := range solution(candidates[i:], target-c) {
			res = append(res, append([]int{c}, sol...))
		}
	}
	return res
}

func quickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	base := nums[0]
	i, j := 0, len(nums) - 1
	for i < j {
		for nums[j] >= base && i<j {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]

		for nums[i] < base && i < j {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	quickSort(nums[:i])
	quickSort(nums[i+1:])
}

func Test_(t *testing.T) {
	candidates := []int{2,3,5}
	target := 8
	fmt.Println(combinationSum(candidates, target))
}
