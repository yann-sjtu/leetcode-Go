package _8_count_and_say

import (
	"fmt"
	"testing"
)

func countAndSay(n int) string {
	s := "1"
	for i:=1;i<n;i++{
		s = say(s)
	}
	return s
}

func say(s string) string {
	var res string
	var count byte
	l := len(s)
	for i:=1;i<l;i++ {
		if s[i] != s[i-1] {
			res += string([]byte{count+'1', s[i-1]})
			count = 0
			continue
		}
		count++
	}
	res += string([]byte{count+'1', s[l-1]})
	return res
}

func Test_(t *testing.T) {
	fmt.Println(countAndSay(8))
}