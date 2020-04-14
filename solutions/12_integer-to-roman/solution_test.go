package _2_integer_to_roman

import (
	"fmt"
	"testing"
)

// 预先申请一个合适大小的[]byte，拼接时不需要申请新的内存，比直接拼接string更高效
func intToRoman(num int) string {
	byteS := make([]byte, 0, 20)
	keys := [][]byte{[]byte("M"), []byte("CM"), []byte("D"), []byte("CD"), []byte("C"), []byte("XC"), []byte("L"), []byte("XL"), []byte("X"), []byte("IX"), []byte("V"), []byte("IV"), []byte("I")}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for i := 0; i < 13; i++ {
		for num >= values[i] {
			num -= values[i]
			byteS = append(byteS, keys[i]...)
		}
	}
	return string(byteS)
}

func Test_(t *testing.T) {
	n := 1994
	fmt.Println(intToRoman(n))
}
