package _6_ThreeSumClosest

func threeSumClosest(nums []int, target int) int {
	l := len(nums)
	if l < 3 {
		return 0
	}
	quickSort(nums)
	diff := nums[0] + nums[1] + nums[2] - target
	for i := range nums[:l-2] {
		if nums[i] > target/3 {
			diff = closestDiff(diff, nums[i]+nums[i+1]+nums[i+2] - target)
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for left, right := i+1, l-1; left < right; {
			if nums[i]+nums[left]+nums[right] == target {
				return target
			} else if nums[i]+nums[left]+nums[right] < target {
				diff = closestDiff(diff, nums[i]+nums[left]+nums[right] - target)
				left++
			} else if nums[i]+nums[left]+nums[right] > target {
				diff = closestDiff(diff, nums[i]+nums[left]+nums[right] - target)
				right--
			}
		}
	}

	return target + diff
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
