package __palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	helper := 1
	for temp := x; temp >= 10; temp /= 10 {
		helper *= 10
	}
	var first, last int
	for helper >= 10 {
		last = x % 10
		first = x / helper
		if first != last {
			return false
		}
		x -= first * helper
		x /= 10
		helper /= 100
	}
	return true
}

func isPalindromeV2(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	var helper, remainder int
	for helper < x {
		remainder = x % 10
		helper = helper*10 + remainder
		x /= 10
	}
	if helper == x || helper/10 == x {
		return true
	}
	return false
}
