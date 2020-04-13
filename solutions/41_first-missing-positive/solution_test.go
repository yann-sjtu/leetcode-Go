package _1_first_missing_positive

import (
	"fmt"
	"testing"
)

func firstMissingPositive(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 1
	}
	for i := l-1;i>=0;i-- {
		for ;nums[i] > 0 && nums[i] <= l && nums[i] != i+1 && nums[i] != nums[nums[i]-1]; {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, n := range nums {
		if n != i+1 {
			return i+1
		}
	}
	return l+1
}

func Test_(T *testing.T) {
	l := []int{3,4,-1,1}
	fmt.Println(firstMissingPositive(l))
}
