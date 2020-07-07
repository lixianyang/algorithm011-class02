package Week_02

// 大顶堆：nlog(k)
func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}
	// 前 k 个元素建堆
	for i := k/2 - 1; i >= 0; i-- {
		heapify(arr, k, i)
	}
	// 遍历其余的数，小于堆顶的加入堆中
	for i := k; i < len(arr); i++ {
		if arr[i] < arr[0] {
			arr[0] = arr[i]
			heapify(arr, k, 0)
		}
	}
	// 堆中保存的就是最小的 k 个数
	return arr[:k]
}
func heapify(arr []int, length, i int) {
	var largest, l, r int
	for {
		largest, l, r = i, 2*i+1, 2*i+2
		if l < length && arr[largest] < arr[l] {
			largest = l
		}
		if r < length && arr[largest] < arr[r] {
			largest = r
		}
		if largest == i {
			break
		}
		arr[largest], arr[i] = arr[i], arr[largest]
		i = largest
	}
}

// 暴力：排序，根据输入数据特性优化
/*
func getLeastNumbers(arr []int, k int) []int {
    counter := make([]int, 10001)
    for _, a := range arr {
        counter[a]++
    }
    result := make([]int, 0, k)
    for i, n := range counter {
        for ; n > 0; n-- {
            result = append(result, i)
            if len(result) == k {
                return result
            }
        }
    }
    return nil
}
*/
// 暴力：排序 nlog(n)
/*
func getLeastNumbers(arr []int, k int) []int {
    sort.Ints(arr)
    return arr[:k]
}
*/
