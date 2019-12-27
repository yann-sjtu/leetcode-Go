package _4_longest_common_prefix

import (
	"fmt"
	"testing"
)

func longestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 0 {
		return ""
	}
	prefix := strs[0]
	for prefixLen := len(prefix); prefixLen > 0; {
		for i, str := range strs {
			if len(str) < prefixLen || str[0:prefixLen] != prefix {
				prefix = prefix[0 : prefixLen-1]
				prefixLen = len(prefix)
				break
			}
			if i == l-1 {
				return prefix
			}
		}
	}

	return prefix
}

func longestCommonPrefixV2(strs []string) string {
	l := len(strs)
	if l == 0 || len(strs[0]) == 0 {
		return ""
	}
	var index int
	for index < len(strs[0]) {
		c := strs[0][index]
		for _, str := range strs[1:] {
			if len(str) < index+1 || str[index] != c {
				return strs[0][:index]
			}
		}
		index++
	}

	return strs[0]
}

func Test_(t *testing.T) {
	s := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefixV2(s))
}
