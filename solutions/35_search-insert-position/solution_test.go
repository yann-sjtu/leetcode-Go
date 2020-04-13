package _5_search_insert_position

import (
	"fmt"
	"testing"
)

func searchInsert(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	var left, right int
	for left, right = 0, l-1; left < right; {
		mid := (left + right) /2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	if nums[left] > target {
		return left - 1
	}
	return left+1
}

func Test_(t *testing.T) {
	nums := []int{1,3,5,6}
	fmt.Println(searchInsert(nums, 5))
}
