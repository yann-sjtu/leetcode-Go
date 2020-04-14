package _7_remove_element

func removeElement(nums []int, val int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	i := 0
	for j := 0; j < l; j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
