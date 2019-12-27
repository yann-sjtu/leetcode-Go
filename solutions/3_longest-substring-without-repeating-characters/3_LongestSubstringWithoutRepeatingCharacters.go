package __longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	var start, max int
	m := make(map[rune]int)
	for index, character := range s {
		if mapIndex, ok := m[character]; !ok || mapIndex < start {
			if index - start + 1 > max {
				max = index - start + 1
			}
		} else {
			start = mapIndex+1
		}
		m[character] = index
	}
	return max
}