func permute(nums []int) [][]int {
    n := len(nums)
    if n == 0 { return nil }
    var result [][]int
    trace := make([]int, 0, n)
    visited := make([]bool, n)   // 记录 nums 下标已经访问过
    backtrack(nums, trace, visited, &result)
    return result
}
func backtrack(nums []int, trace []int, visited []bool, result *[][]int) {
    if len(nums) == len(trace) {
        dst := make([]int, len(trace))
        copy(dst, trace)
        *result = append(*result, dst)
        return
    }
    for i := range nums {
        // 将选择从选择列表里移除
        if visited[i] { continue }
        visited[i] = true
        trace = append(trace, nums[i])
        backtrack(nums, trace, visited, result)
        trace = trace[:len(trace)-1]
        visited[i] = false
    }

    return
}
