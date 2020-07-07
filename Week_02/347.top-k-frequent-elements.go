package Week_02

// 小顶堆解决最大 k 问题
func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)
	for i := range nums {
		counter[nums[i]]++
	}
	n := len(counter)
	items := make([][2]int, 0, n)
	for num, count := range counter {
		items = append(items, [2]int{count, num})
	}
	// 前 k 个元素建堆
	for i := k/2 - 1; i >= 0; i-- {
		heapify(items, k, i)
	}
	// 小元素都移出去，剩下的就是大元素
	for i := k; i < n; i++ {
		if items[i][0] > items[0][0] {
			items[0] = items[i]
			heapify(items, k, 0)
		}
	}
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = items[i][1]
	}
	return result
}
func heapify(arr [][2]int, length, i int) {
	var min, l, r int
	for {
		min, l, r = i, 2*i+1, 2*i+2
		if l < length && arr[l][0] < arr[min][0] {
			min = l
		}
		if r < length && arr[r][0] < arr[min][0] {
			min = r
		}
		if min == i {
			break
		}
		arr[min], arr[i] = arr[i], arr[min]
		i = min
	}
}
