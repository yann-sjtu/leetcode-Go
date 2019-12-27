package _0_valid_parentheses

func isValid(s string) bool {
	l := len(s)
	if l==0 {
		return true
	}
	if l&1 == 1 {
		return false
	}
	arr := make([]byte, 0, l)
	arr = append(arr, 0)
	for i:=0;i<l;i++{
		if isRight(s[i]) {
			last := len(arr) - 1
			left := arr[last]
			if s[i] - left > 2 {
				return false
			}
			arr = arr[:last]
		} else {
			arr = append(arr, s[i])
		}
	}
	if len(arr) > 1 {
		return false
	}
	return true
}

func isRight(b byte) bool {
	if b == ')' || b==']' || b=='}' {
		return true
	}
	return false
}
