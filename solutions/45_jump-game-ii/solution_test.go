package _5_jump_game_ii

import (
	"fmt"
	"testing"
)

func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	var count int
	for i:=0; i<len(nums); {
		if i+nums[i] >= len(nums)-1 {
			count++
			return count
		}
		var next, max int
		for j:=i+1; j<len(nums) && j<=i+nums[i]; j++ {
			if j+nums[j] > max {
				max = j+nums[j]
				next = j
			}
		}
		i = next
		count++
	}
	return count
}

func Test_(t *testing.T) {
	fmt.Println(jump([]int{1,1,1,1}))
}
