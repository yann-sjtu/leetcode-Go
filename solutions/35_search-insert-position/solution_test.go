package _5_search_insert_position

func searchInsert(nums []int, target int) int {
	l := len(nums)
	if l == 0 || target < nums[0] {
		return 0
	}
	var left, right int
	for left, right = 0, l; left < right; {
		mid := (left + right) /2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else {
			return mid
		}
	}
	return right
}
