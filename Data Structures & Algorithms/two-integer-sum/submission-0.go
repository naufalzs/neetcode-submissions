func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
	for i, v := range nums {
		if j, ok := seen[v];ok {
			return []int{j, i}
		}
		seen[target - v] = i
	}
	return nums
}
