package _5_search_insert_position

func searchInsert(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	for left, right := 0, l; left < right; {
		mid := (left + right) /2
		if nums[mid] >
	}
}
