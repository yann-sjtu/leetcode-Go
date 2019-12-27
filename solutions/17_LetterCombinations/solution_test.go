package _7_LetterCombinations

import (
	"fmt"
	"testing"
)

func letterCombinations(digits string) []string {
	l := len(digits)
	if l == 0 {
		return nil
	}
	m := make(map[byte]string)
	m['2'] = "abc"
	m['3'] = "def"
	m['4'] = "ghi"
	m['5'] = "jkl"
	m['6'] = "mno"
	m['7'] = "pqrs"
	m['8'] = "tuv"
	m['9'] = "wxyz"

	res := []string{""}
	for i := 0; i < l; i++ {
		var newRes []string
		for _, s := range res {
			for _, c := range m[digits[i]] {
				newRes = append(newRes, s+string(c))
			}
		}
		res = newRes
	}

	return res
}

func Test_(t *testing.T) {
	res := letterCombinations("23")
	fmt.Println(res)
}
