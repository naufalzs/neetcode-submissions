func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	// map it into bucket sort, knowing that the frequencies not gonna exceed length of nums
	numsLen := len(nums) + 1
	bucket := make([][]int, numsLen)
	for num, freq := range count {
		bucket[freq] = append(bucket[freq], num)
	}

	// take top 2 freq
	res := make([]int, 0, k)
	for i := len(bucket) - 1; i >= 0; i-- {
		for _, num := range bucket[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
