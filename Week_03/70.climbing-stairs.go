package Week_03

// 循环
func climbStairs(n int) int {
	p1, p2 := 0, 1
	for ; n > 0; n-- {
		p1, p2 = p2, p1+p2
	}
	return p2
}

// 递归加缓存
/*
func climbStairs(n int) int {
    h := map[int]int{1: 1, 2: 2, 3: 3}
    var f func(int) int
    f = func(n int) int {
        if v, ok := h[n]; ok {
            return v
        }
        h[n] = f(n-1)+f(n-2)
        return h[n]
    }
    return f(n)
}
*/
// 递归 f(n)=f(n-1)+f(n-2) 超时
/*
func climbStairs(n int) int {
    if n <= 3 {
        return n
    }
    return climbStairs(n-1)+climbStairs(n-2)
}
*/
