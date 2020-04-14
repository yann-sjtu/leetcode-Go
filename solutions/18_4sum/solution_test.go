package _8_4sum

func fourSum(nums []int, target int) [][]int {
	l := len(nums)
	if l < 4 {
		return nil
	}
	quickSort(nums)
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
