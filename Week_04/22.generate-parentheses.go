func generateParenthesis(n int) []string {
    if n == 0 { return nil }
    result := make([]string, 0)
    trace := make([]byte, 0, n)
    backtrack(trace, 0, 0, n, &result)
    return result
}
func backtrack(trace []byte, left, right, n int, result *[]string) {
    if left == n && right == n {
        *result = append(*result, string(trace))
        return
    }
    if left > n || right > left { return }
    backtrack(append(trace, '('), left+1, right, n, result)
    backtrack(append(trace, ')'), left, right+1, n, result)
    return
}
