package _6_permutations

import (
	"fmt"
	"testing"
)

// time cost 0ms, defeat 100%
// memory cost 2.8MB, defeat 100%
// dfs递归，每次递归要开辟新的内存。本打算先随便写个ac了之后再整理，结果发现leetcode应该是不会重复计算新申请的内存，只会计算gc之后还占用的内存的最大值
func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	dfs(&ans,nil, nums)
	return ans
}

func dfs(ans *[][]int, trace, choice []int) {
	if len(choice) == 0 {
		*ans = append(*ans, trace)
		return
	}
	for i := range choice {
		nextTrace := make([]int, 0, len(trace)+1)
		nextTrace = append(nextTrace, trace...)
		nextTrace = append(nextTrace, choice[i])
		newChoice := make([]int, 0, len(choice)-1)
		newChoice = append(newChoice, choice[:i]...)
		newChoice = append(newChoice, choice[i+1:]...)
		dfs(ans, nextTrace, newChoice)
	}
}


func Test_(t *testing.T) {
	fmt.Println(permute([]int{1,2,3}))
}
