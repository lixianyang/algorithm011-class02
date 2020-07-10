// 分治
func myPow(x float64, n int) float64 {
    if n < 0 {
        x = 1 / x
        n = -n
    }
    return pow(x, n)
}
func pow(x float64, n int) float64 {
    if n == 0 { return 1 }
    half := pow(x, n/2)
    if n % 2 == 1 {
        return half*half*x
    }
    return half*half
}
// 迭代
func myPow(x float64, n int) float64 {
    if n == 0 { return 1 }
    if n < 0 {
        n = -n
        x = 1/x
    }
    result := 1.0
    for n > 0 {
        if n & 1 == 1 {
            result *= x
        }
        x *= x
        n = n >> 1
    }
    return result
}
