package _2_generate_parentheses

//dynamic programming
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	if n == 1 {
		return []string{"()"}
	}
	var res []string
	for i:=0;i<=n-1;i++{
		for _, s1 := range generateParenthesis(n-1-i) {
			for _, s2 := range generateParenthesis(i) {
				res = append(res, "(" + s1 + ")"+s2)
			}
		}
	}
	return res
}
