package _1_next_permutation

import (
	"fmt"
	"testing"
)

func nextPermutation(nums []int)  {
	l := len(nums)
	if l < 2 {
		return
	}
	var i int
	for i=l-1;i>0;i-- {
		if nums[i-1] < nums[i] {
			break
		}
	}
	if i == 0 {
		reverse(nums)
		return
	}
	if i == l-1 {
		nums[l-2], nums[l-1] = nums[l-1], nums[l-2]
		return
	}
	for j := l-1;j>i-1;j-- {
		if nums[j] > nums[i-1] {
			nums[i-1], nums[j] = nums[j], nums[i-1]
			break
		}
	}
	reverse(nums[i:])
}

func reverse(nums []int) {
	l := len(nums)
	for i:=0;i<l/2;i++ {
		nums[i], nums[l-i-1] = nums[l-i-1], nums[i]
	}
}

func Test_(t *testing.T) {
	nums := []int{1, 3, 2}
	nextPermutation(nums)
	fmt.Println(nums)
}
