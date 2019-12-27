package _3_RomanToInteger

import (
	"fmt"
	"testing"
)

// map比slice占空间大
func romanToInt(s string) int {
	//keys := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	//values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	m := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}
	var num int
	for len(s) > 1 {
		if v, ok := m[s[0:2]]; ok {
			num += v
			s = s[2:]
		} else if v, ok := m[s[0:1]]; ok {
			num += v
			s = s[1:]
		}
	}
	if len(s) != 0 {
		num += m[s]
	}
	return num
}

func Test_(t *testing.T) {
	s := "IV"
	fmt.Println(romanToInt(s))
}
