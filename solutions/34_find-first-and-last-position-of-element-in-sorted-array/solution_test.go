package _4_find_first_and_last_position_of_element_in_sorted_array

import (
	"fmt"
	"testing"
)

func searchRange(nums []int, target int) []int {
	res := make([]int, 0, 2)
	res = append(res, leftBound(nums, target))
	if res[0] == -1 {
		res = append(res, -1)
	} else {
		res = append(res, rightBound(nums, target))
	}
	return res
}

func leftBound(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	var start, end int
	for start, end = 0, l; start < end; {
		mid := (start + end) / 2
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	if start == l || nums[start] != target {
		return -1
	}
	return start
}

func rightBound(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	var start, end int
	for start, end = 0, l; start < end; {
		mid := (start + end) / 2
		if nums[mid] > target {
			end = mid
		} else {
			start = mid + 1
		}
	}
	if start == 0 || nums[start-1] != target {
		return -1
	}
	return start-1
}

func Test_(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(nums, 8))
}
