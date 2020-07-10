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
