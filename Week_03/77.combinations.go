package Week_03

// 递归：f(n,k) = f(n-1,k-1)+f(n-1,k)
// f(n-1,k-1) 表示 n-1 个数中所有 k-1 个数的组合，最后一个数是 n
// f(n-1,k) 表示 n-1 个数中所有 k 个数的组合（不包含 n ）
func combine(n int, k int) [][]int {
	if n*k == 0 {
		return nil
	}
	result := make([][]int, 0, n)
	if k == 1 {
		for i := 1; i <= n; i++ {
			result = append(result, []int{i})
		}
		return result
	}
	result = append(result, combine(n-1, k)...)
	for _, item := range combine(n-1, k-1) {
		result = append(result, append(item, n))
	}
	return result
}
