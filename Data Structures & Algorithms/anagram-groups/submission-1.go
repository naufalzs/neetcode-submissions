func groupAnagrams(strs []string) [][]string {
	temp := [][26]int{}
	for _, str := range strs {
		arr := [26]int{}
		for _, ch := range str {
			arr[ch - 'a']++
		}
		temp = append(temp, arr)
	}

	res := [][]string{}
	seen := make(map[[26]int]int)
	for i, v := range temp {
		idx, ok := seen[v]
		if !ok {
			idx = len(res)
			seen[v] = idx
			res = append(res, []string{})
		}
		res[idx] = append(res[idx], strs[i])
	}
	return res
}
