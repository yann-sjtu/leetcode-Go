package solutions

import (
	"fmt"
	"testing"
)

func threeSum(nums []int) [][]int {
	l := len(nums)
	if l < 3 {
		return nil
	}
	quickSort(nums)
	var res [][]int
	for i := range nums[:l-2] {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for left, right := i+1, l-1; left < right; {
			if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for right > left && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}
	return res
}

func quickSort(nums []int) {
	l := len(nums)
	if l < 2 {
		return
	}

	head, tail := 1, l-1
	privot := nums[0]
	for head < tail {
		for head < tail && nums[head] < privot {
			head++
		}
		for tail > head && nums[tail] >= privot {
			tail--
		}
		nums[head], nums[tail] = nums[tail], nums[head]
	}
	if nums[tail] < privot {
		nums[0], nums[tail] = nums[tail], nums[0]
	}

	quickSort(nums[:head])
	quickSort(nums[tail:])
}

func threeSumClosest(nums []int, target int) int {
	l := len(nums)
	if l < 3 {
		return 0
	}
	quickSort(nums)
	diff := nums[0] + nums[1] + nums[2] - target
	for i := range nums[:l-2] {
		if nums[i] > target/3 {
			diff = closestDiff(diff, nums[i]+nums[i+1]+nums[i+2]-target)
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for left, right := i+1, l-1; left < right; {
			if nums[i]+nums[left]+nums[right] == target {
				return target
			} else if nums[i]+nums[left]+nums[right] < target {
				diff = closestDiff(diff, nums[i]+nums[left]+nums[right]-target)
				left++
			} else if nums[i]+nums[left]+nums[right] > target {
				diff = closestDiff(diff, nums[i]+nums[left]+nums[right]-target)
				right--
			}
		}
	}

	return target + diff
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func closestDiff(diff1, diff2 int) int {
	if abs(diff1) < abs(diff2) {
		return diff1
	}
	return diff2
}

func fourSum(nums []int, target int) [][]int {
	l := len(nums)
	if l < 4 {
		return nil
	}
	quickSort(nums)
	fmt.Println(nums)
	var res [][]int
	for index1 := range nums[:l-3] {
		if index1 > 0 && nums[index1] == nums[index1-1] {
			continue
		}
		for index2 := index1 + 1; index2 < l-2; index2++ {
			if index2 > index1+1 && nums[index2] == nums[index2-1] {
				continue
			}
			for left, right := index2+1, l-1; left < right; {
				if nums[index1]+nums[index2]+nums[left]+nums[right] == target {
					res = append(res, []int{nums[index1], nums[index2], nums[left], nums[right]})
					left++
					right--
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				} else if nums[index1]+nums[index2]+nums[left]+nums[right] < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return res
}

func isValid(s string) bool {
	l := len(s)
	if l==0 {
		return true
	}
	if l&1 == 1 {
		return false
	}
	arr := make([]byte, 0, l/2)
	arr = append(arr, 0)
	for i:=0;i<l;i++{
		if isRight(s[i]) {
			last := len(arr) - 1
			left := arr[last]
			if s[i] - left > 2 {
				return false
			}
			arr = arr[:last]
		} else {
			arr = append(arr, s[i])
		}
	}
	if len(arr) > 1 {
		return false
	}
	return true
}

func isRight(b byte) bool {
	if b == ')' || b==']' || b=='}' {
		return true
	}
	return false
}

func Test_(t *testing.T) {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(isValid("()[]{}"))
}
