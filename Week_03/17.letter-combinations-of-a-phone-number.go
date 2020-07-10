func letterCombinations(digits string) []string {
    if len(digits) == 0 { return nil }
    m := map[byte][]byte{
        '2': []byte{'a', 'b', 'c'},
        '3': []byte{'d', 'e', 'f'},
        '4': []byte{'g', 'h', 'i'},
        '5': []byte{'j', 'k', 'l'},
        '6': []byte{'m', 'n', 'o'},
        '7': []byte{'p', 'q', 'r', 's'},
        '8': []byte{'t', 'u', 'v'},
        '9': []byte{'w', 'x', 'y', 'z'},
    }
    result := make([]string, 0)
    trace := make([]byte, 0, len(digits))
    backtrack(&result, []byte(digits), trace, m)
    return result
}
func backtrack(result *[]string, digits []byte, trace []byte, m map[byte][]byte) {
    if len(digits) == 0 {
        *result = append(*result, string(trace))
        return
    }
    for _, c := range m[digits[0]] {
        backtrack(result, digits[1:], append(trace, c), m)
    }
}
