package Week_02

func twoSum(nums []int, target int) []int {
	h := make(map[int]int)
	for i := range nums {
		if j, ok := h[target-nums[i]]; ok {
			return []int{j, i}
		}
		h[nums[i]] = i
	}
	return nil
}
