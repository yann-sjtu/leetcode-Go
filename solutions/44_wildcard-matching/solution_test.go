package _4_wildcard_matching

import (
	"fmt"
	"testing"
)

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i:=0;i<=len(s);i++{
		dp[i] = make([]bool, len(p)+1)
		dp[i][0] = false
	}
	dp[0][0] = true
	for i:=1;i<=len(p);i++{
		dp[0][i] = dp[0][i-1] && p[i-1] == '*'
	}
	for i:=1;i<=len(s);i++ {
		for j:=1;j<=len(p);j++{
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			}
		}
	}

	return dp[len(s)][len(p)]
}

//双指针贪心算法
// time cost 0ms, defeated 100%
// memory cost 3.1MB, defeat 100%
func isMatchV2(s string, p string) bool {
	sStar, pStar := -1, -1
	var si, pi int
	sl, pl := len(s), len(p)
	for si < sl {
		if pi < pl && (s[si] == p[pi] || p[pi] == '?') {
			si++
			pi++
		} else if pi < pl && p[pi] == '*' {
			sStar, pStar = si, pi
			pi++
		} else if pStar != -1 {
			pi = pStar+1
			sStar++
			si = sStar
		} else {
			return false
		}
	}

	var end int
	for end=pi;end<pl && p[end]=='*';end++ {}

	return end == pl
}


func Test_(t *testing.T) {
	fmt.Println(isMatchV2("adceb", "*a*b"))
}