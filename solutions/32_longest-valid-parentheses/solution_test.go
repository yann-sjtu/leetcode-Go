package _2_longest_valid_parentheses

import (
	"fmt"
	"testing"
)

func longestValidParentheses(s string) int {
	c := byte('(')
	l := len(s)
	var left, right, count int
	for i:=0;i<l;i++ {
		if s[i] == c {
			left++
		}else {
			right++
		}
		if left == right {
			if 2*left > count {
				count = 2*left
			}
		} else if right > left {
			left, right = 0,0
		}
	}
	left, right = 0, 0
	for i:=l-1;i>=0;i-- {
		if s[i] == c {
			left++
		}else {
			right++
		}
		if left == right {
			if 2*left > count {
				count = 2*left
			}
		} else if right < left {
			left, right = 0,0
		}
	}
	return count
}

func Test_(t *testing.T) {
	s := "(()"
	fmt.Println(longestValidParentheses(s))
}