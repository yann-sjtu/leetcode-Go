package _8_implement_strstr

import (
	"fmt"
	"testing"
)

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

//KMP algorithm
func strStrV2(haystack string, needle string) int {
	l1, l2 := len(haystack), len(needle)
	if l2 == 0 {
		return 0
	}
	if l1 < l2 {
		return -1
	}
	Next := getNext(needle)
	for i, j := 0, 0; i < l1 && j < l2; i++ {
		if haystack[i] == needle[j] {
			if j == l2-1 {
				return i - j
			}
			j++
		} else if j != 0 {
			j = Next[j-1]
			i--
		}
	}

	return -1
}

func getNext(s string) []int {
	l := len(s)
	if l < 1 {
		return nil
	}
	arr := make([]int, l)
	arr[0] = 0
	i, j := 0, 1
	for j < l {
		if s[i] == s[j] {
			arr[j] = i + 1
			i++
			j++
		} else {
			if i == 0 {
				arr[j] = 0
				j++
			} else {
				i = arr[i-1]
			}
		}
	}
	return arr
}

func Test_getNext(t *testing.T) {
	s := "issip"
	fmt.Println(getNext(s))
}

func Test_str(t *testing.T) {
	s1 := "mississippi"
	s2 := "issip"
	fmt.Println(strStr(s1, s2))
	fmt.Println(strStrV2(s1, s2))
}
