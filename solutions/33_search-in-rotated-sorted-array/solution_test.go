package _3_search_in_rotated_sorted_array

import (
	"fmt"
	"testing"
)

func search(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	var mid int
	for start, end := 0, l-1; start<=end; {
		mid = (start+end)/2
		if target == nums[mid] {
			return mid
		}
		//first half sorted
		if nums[mid] < nums[end] {
			if target > nums[mid] && target <= nums[end] {
				start = mid+1
			} else {
				end = mid-1
			}
		} else { //second half sorted
			if target >= nums[start] && target < nums[mid] {
				end = mid-1
			} else {
				start = mid+1
			}
		}
	}
	return -1
}

func binarySearch(nums []int, target int) int {
	l := len(nums)
	for start, end := 0, l-1; start<=end; {
		mid := (start+end)/2
		if nums[mid] < target {
			start = mid+1
		} else if nums[mid] > target{
			end = mid-1
		} else {
			return mid
		}
	}
	return -1
}


func Test_(t *testing.T) {
	nums := []int{3,4,5,1,2}
	fmt.Println(search(nums, 2))

}