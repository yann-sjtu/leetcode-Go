package _5_ThreeSum

import (
	"fmt"
	"testing"
)

func threeSum(nums []int) [][]int {
	l := len(nums)
	if l < 3 {
		return nil
	}
	m := make(map[int]struct{})
	res := make([][]int, 0)
	for i := 0; i < l; i++ {
		if _, ok := m[nums[i]]; ok {
			continue
		}
		m[nums[i]] = struct{}{}
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					res = append(res, sort(nums[i], nums[j], nums[k]))
				}
			}
		}
	}
	return res
}

func sort(num1, num2, num3 int) []int {
	nums := make([]int, 0, 3)
	if num1 < num2 {
		nums = append(nums, num1, num2)
	} else {
		nums = append(nums, num2, num1)
	}

	if num3 < num1 {
		nums = append([]int{num3}, nums...)
	} else if num3 > num2 {
		nums = append(nums, num3)
	} else {
		nums = []int{num1, num3, num2}
	}
	return nums
}

func Test_(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println(res)
}
