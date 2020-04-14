package _2_trapping_rain_water

import (
	"fmt"
	"testing"
)

//按高度遍历，效率较低
func trap(height []int) int {
	var result int
	for h:=1;;h++ {
		start, end := -1, -1
		var count int
		for i, bar := range height {
			if bar >= h {
				if start == -1 {
					start = i
				} else {
					end = i
				}
				count++
			}
		}
		if end == -1 {
			break
		}
		result += end-start-count+1
	}
	return result
}

//按列遍历，dp解法
func trapV2(height []int) int {
	var sum int
	l := len(height)
	if l < 3 {
		return sum
	}
	maxLeft, maxRight := make([]int, l), make([]int, l)
	for i:=1;i<l;i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i-1])
	}
	for i:=l-2;i>=0;i-- {
		maxRight[i] = max(maxRight[i+1], height[i+1])
	}
	for i, h := range height {
		if h < min(maxLeft[i], maxRight[i]) {
			sum += min(maxLeft[i], maxRight[i]) - h
		}
	}

	return sum
}

//按列遍历,用单变量代替maxLeft数组
func trapV3(height []int) int {
	var sum int
	l := len(height)
	if l < 3 {
		return sum
	}
	maxRight := make([]int, l)
	for i:=l-2;i>=0;i-- {
		maxRight[i] = max(maxRight[i+1], height[i+1])
	}
	var maxLeft int
	for i, h := range height {
		if i>0 {
			maxLeft = max(maxLeft, height[i-1])
		}
		if h < min(maxLeft, maxRight[i]) {
			sum += min(maxLeft, maxRight[i]) - h
		}
	}

	return sum
}

//按列遍历,双指针法双向遍历
func trapV4(height []int) int {
	var sum int
	l := len(height)
	if l < 3 {
		return sum
	}
	var maxLeft, maxRight int
	left, right := 1, l-2
	for left <= right {
		if height[left-1] < height[right+1] {
			maxLeft = max(maxLeft, height[left-1])
			if maxLeft > height[left] {
				sum += maxLeft - height[left]
			}
			left++
		} else {
			maxRight = max(maxRight, height[right+1])
			if maxRight > height[right] {
				sum += maxRight - height[right]
			}
			right--
		}
	}

	return sum
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func min(n1, n2 int) int {
	if n1 > n2 {
		return n2
	}
	return n1
}

type Stack struct {
	elements []int
}

func (s *Stack) Push(n int) {
	s.elements = append(s.elements, n)
}

func (s *Stack) Pop() int {
	l := len(s.elements)
	if l == 0 {
		panic("empty")
	}
	element := s.elements[l-1]
	s.elements = s.elements[:l-1]
	return element
}

func (s *Stack) Peek() int {
	l := len(s.elements)
	if l == 0 {
		panic("empty")
	}
	return s.elements[l-1]
}

func (s *Stack) Len () int {
	return len(s.elements)
}

//用栈实现
func trapV5(height []int) int {
	var sum int
	l := len(height)
	if l < 3 {
		return sum
	}
	stack := &Stack{}
	for i := range height {
		for stack.Len() != 0 && height[i] > height[stack.Peek()] {
			lastIndex := stack.Pop()
			if stack.Len() == 0 {
				break
			}
			sum += (min(height[i], height[stack.Peek()]) - height[lastIndex]) * (i-stack.Peek()-1)
		}
		stack.Push(i)
	}
	return sum
}


func Test_(t *testing.T) {
	bars := []int{4,2,0,3,2,5}
	fmt.Println(trapV5(bars))
}