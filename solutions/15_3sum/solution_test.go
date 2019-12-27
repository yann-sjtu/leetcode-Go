package _5_3sum

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
