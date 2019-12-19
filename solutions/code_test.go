package solutions

import (
	"fmt"
	"testing"
)

func TestLeetcode(t *testing.T) {
	s := "cb1bc"
	res := isPalindrome(s)
	fmt.Println(res)
	res2 := longestPalindrome(s)
	fmt.Println(res2)
}
