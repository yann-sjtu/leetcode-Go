package __two_sum

//a really brute force way
func twoSum(nums []int, target int) []int {
	for i, v1 := range nums {
		for j, v2 := range nums[i+1:] {
			if v1 + v2 == target {
				return []int{i, j+i+1}
			}
		}
	}
	return nil
}

//use addition hash map
func twoSumV2(nums []int, target int) []int {
	m := make(map[int]int) //key=num, value=index
	for i, v1 := range nums {
		if j, ok := m[target-v1]; ok {
			return []int{j, i}
		}
		m[v1]=i
	}
	return nil
}

