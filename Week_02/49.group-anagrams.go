package Week_02

// 这里字典的 key 是 26 位数组，比较时间比较满，可以考虑优化
func groupAnagrams(strs []string) [][]string {
	const n = 26
	h := make(map[[n]int][]string)
	for i := range strs {
		var counter [n]int
		for _, c := range strs[i] {
			counter[c-'a']++
		}
		h[counter] = append(h[counter], strs[i])
	}

	result := make([][]string, 0, len(h))
	for _, v := range h {
		result = append(result, v)
	}

	return result
}
