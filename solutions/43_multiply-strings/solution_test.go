package _3_multiply_strings

import (
	"fmt"
	"testing"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l1, l2 := len(num1), len(num2)
	result := make([]byte, l1+l2)
	for i:=l1-1;i>=0;i-- {
		n1 := num1[i] - '0'
		for j:=l2-1;j>=0;j-- {
			n2 := num2[j] - '0'
			resIndex := i+j+1
			//高位result[resIndex-1]此时可能是两位数，不过不必急于处理进位，for j执行到下个循环时会一起处理进位
			result[resIndex-1] += (result[resIndex] + n1*n2) / 10
			result[resIndex] = (result[resIndex] + n1*n2) % 10
		}
	}

	//start := -1
	//for i, c := range result {
	//	if start == -1 && c != 0 {
	//		start = i
	//	}
	//	result[i] += '0'
	//}
	//拆开用for循环，比在for循环里加判断语句执行效率高
	var start int
	for ;start<l1+l2 && result[start] == 0;start++ {}
	for i:=start;i<l1+l2;i++{
		result[i] +='0'
	}

	return string(result[start:])
}

func Test_(t *testing.T) {
	//56088
	fmt.Println(multiply("123", "456"))
}
