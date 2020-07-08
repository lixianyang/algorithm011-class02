package Week_03

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	generate(0, 0, n, "", &result)
	return result
}
func generate(left, right, n int, s string, result *[]string) {
	if left == n && right == n {
		*result = append(*result, s)
		return
	}
	if left < n {
		generate(left+1, right, n, s+"(", result)
	}
	if right < left {
		generate(left, right+1, n, s+")", result)
	}
}
