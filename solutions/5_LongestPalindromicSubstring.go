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

//动态规划法
func longestPalindromeV2(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}
	var max int
	palindrome := s[0:1]
	P := make([][]bool, l)
	for i:= range P {
		P[i] = make([]bool, l)
		P[i][i] = true
	}
	for i:= range P[1:] {
		P[i][i+1] = s[i] == s[i+1]
		if P[i][i+1] {
			max = 2
			palindrome = s[i:i+2]
		}
	}
	//待优化，可以提前退出
	for interval := 2; interval < len(s); interval++ {
		for i:=0;i<l-interval;i++ {
			P[i][i+interval] = P[i+1][i+interval-1] && (s[i] == s[i+interval])
			if P[i][i+interval] && interval + 1 > max {
				max = interval + 1
				palindrome = s[i:i+interval+1]
			}
		}
	}
	return palindrome
}

// Manacher algorithm
func longestPalindromeV3(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}
	s = preProcess(s)

	P := make([]int, 2*l+3)
	Center := 1
	Right := 1
	MaxR := 0
	P[0] = 1
	P[1] = 1
	for i:=2;i<=2*l+1;i++ {
		mirror_i := 2 * Center - i
		if mirror_i <= 0 {
			for r:=0;i-r>0&&i+r<2*l+2;r++ {
				if s[i-r] != s[i+r] {
					P[i] = r-1
				}
				if r
			}
		}
		if mirror_i > 0 && P[mirror_i] < Right - i {
			P[i] = P[mirror_i]
		}


		for r:=1;i-r>=0&&i+r<2*l+3;r++{
			if s[i-r] == s[i+r] {
				P[i] = 2*r+1
			}
		}
	}



}

func preProcess(s string) string {
	l := len(s)
	startSymbol := []byte("^")
	endSymbol := []byte("$")
	intervalSymbol := []byte("#")
	ss := make([]byte, 0, 2*l+3)
	ss = append(ss, startSymbol...)
	for i:=0;i<2*l+1;i++{
		if i&1==1 {
			ss = append(ss, s[i/2])
		} else {
			ss = append(ss, intervalSymbol...)
		}
	}
	ss = append(ss, endSymbol...)
	return string(ss)
}

