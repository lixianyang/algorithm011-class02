package main

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	j := 0 // 指向无重复元素数组最后一个元素
	for i := 1; i < n; i++ {
		if nums[i-1] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}
