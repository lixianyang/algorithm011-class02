package Week_02

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	deque, result := make([]int, 0, k), make([]int, 0, n-k+1)
	for i := 0; i < n; i++ {
		// 大于时间窗口的元素下标从队头出队
		for len(deque) > 0 && i-deque[0] == k { // i=k 时，队列里最多有k+1个元素，这时需要移除队首元素
			deque = deque[1:]
		}
		// 队尾小于新元素的元素下标从队尾出队
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}
		// 新元素下标入队
		deque = append(deque, i)
		// 队头的元素即为滑动窗口里的最大元素
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}
	return result
}
