// 回溯
func letterCombinations(digits string) []string {
    if len(digits) == 0 { return nil }
    m := [][]byte {
        []byte{'a', 'b', 'c'},
        []byte{'d', 'e', 'f'},
        []byte{'g', 'h', 'i'},
        []byte{'j', 'k', 'l'},
        []byte{'m', 'n', 'o'},
        []byte{'p', 'q', 'r', 's'},
        []byte{'t', 'u', 'v'},
        []byte{'w', 'x', 'y', 'z'},
    }
    result := make([]string, 0)
    trace := make([]byte, 0, len(digits))
    backtrack(&result, []byte(digits), trace, m)
    return result
}
func backtrack(result *[]string, digits []byte, trace []byte, m [][]byte) {
    if len(digits) == 0 {
        *result = append(*result, string(trace))
        return
    }
    for _, c := range m[digits[0]-'2'] {
        backtrack(result, digits[1:], append(trace, c), m)
    }
}
// BFS 记录所有路径作为结果
/*
func letterCombinations(digits string) []string {
    n := len(digits)
    if n == 0 { return nil }
    m := [][]byte{
        []byte{'a', 'b', 'c'},
        []byte{'d', 'e', 'f'},
        []byte{'g', 'h', 'i'},
        []byte{'j', 'k', 'l'},
        []byte{'m', 'n', 'o'},
        []byte{'p', 'q', 'r', 's'},
        []byte{'t', 'u', 'v'},
        []byte{'w', 'x', 'y', 'z'},
    }
    queue := [][]byte{ []byte{} }
    for _, digit := range digits {
        for i := len(queue); i > 0; i-- {
            top := queue[0]
            queue = queue[1:]
            length := len(top)+1
            for _, c := range m[digit-'2'] {
                tmp := make([]byte, 0, length)
                tmp = append(append(tmp, top...), c)
                queue = append(queue, tmp)
            }
        }
    }
    result := make([]string, len(queue))
    for i, arr := range queue {
        result[i] = string(arr)
    }
    return result
}
*/
