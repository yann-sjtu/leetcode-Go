package _1_ContainerWithMostWater

func maxArea(height []int) int {
	l := len(height)
	left := 0
	right := l - 1
	var area int
	for left < right {
		if newArea := min(height[left], height[right]) * (right - left); newArea > area {
			area = newArea
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return area
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
