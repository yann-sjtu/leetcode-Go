package solutions

//中心扩展法
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	var palindrome string
	var max int
	for i :=0; i<len(s);i++ {
		subS := findPalindrome(s, i, false)
		if l:=len(subS);l > max {
			palindrome = subS
			max = l
		}
	}
	for i :=0; i<len(s)-1;i++ {
		subS := findPalindrome(s, i, true)
		if l:=len(subS);l > max {
			palindrome = subS
			max = l
		}
	}


	return palindrome
}

func findPalindrome(s string, center1 int, isEven bool) string {
	center2 := center1
	if isEven {
		center2++
	}
	if s[center1] != s[center2] {
		return ""
	}
	var i int
	for i = 1; center1-i>=0 && center2+i<len(s);i++ {
		if s[center1-i]!=s[center2+i] {
			return s[center1-i+1:center2+i]
		}
	}
	return s[center1-i+1:center2+i]
}