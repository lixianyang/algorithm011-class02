// 计数法
/*
func majorityElement(nums []int) int {
    n := len(nums) / 2
    counter := make(map[int]int, n+1)
    for i := range nums {
        counter[nums[i]]++
        if counter[nums[i]] > n {
            return nums[i]
        }
    }
    return 0
}
*/
// 投票法，最坏情况下一个众数抵消一个非众数，最后剩下的还是众数
func majorityElement(nums []int) int {
    var candidate, count int
    for i := 0; i < len(nums); i++ {
        if count == 0 {
            candidate = nums[i]
        }
        if nums[i] == candidate {
            count++
        } else {
            count--
        }
    }
    return candidate
}
