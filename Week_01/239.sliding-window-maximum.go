package main

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	result, deque := make([]int, 0, n-k+1), make([]int, 0, k)
	for i := 0; i < n; i++ {
		if len(deque) != 0 && deque[0] == i-k {
			deque = deque[1:]
		}
		for len(deque) != 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}
	return result
}
