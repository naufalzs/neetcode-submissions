func hasDuplicate(nums []int) bool {
    amount := make(map[int]int)
	for _, v := range nums {
		if _, ok := amount[v]; ok {
			return true
		} else {
			amount[v] = 1
		}
	}

	return false    
}
