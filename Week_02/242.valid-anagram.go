package Week_02

// 记数法：使用哈希表统计字符串字符数
/*
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    h := make(map[byte]int, len(s))

    for i := range s {
        h[s[i]]++
        h[t[i]]--
    }

    for _, v := range h {
        if v != 0 {
            return false
        }
    }

    return true
}
*/
// 记数法：使用数组优化
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	arr := make([]int, 128)
	for i := range s {
		arr[s[i]]++
		arr[t[i]]--
	}

	for i := range arr {
		if arr[i] != 0 {
			return false
		}
	}

	return true
}

// 暴力法：使用内置排序算法，排序后比较两个字符串是否相同
/*
type Runes []rune

func (r Runes) Len() int { return len(r) }
func (r Runes) Less(i, j int) bool { return r[i] < r[j] }
func (r Runes) Swap(i, j int) { r[i], r[j] = r[j], r[i] }

func SortRunes(data []rune) {
    sort.Sort(Runes(data))
}

func isAnagram(s string, t string) bool {
    l1, l2 := len(s), len(t)
    if l1 != l2 {
        return false
    }
    arr1, arr2 := []rune(s), []rune(t)
    SortRunes(arr1)
    SortRunes(arr2)
    for i := 0; i < l1; i++ {
        if arr1[i] != arr2[i] {
            return false
        }
    }
    return true
}
*/
