import "slices"

func groupAnagrams(strs []string) [][]string {
	tempStr := slices.Clone(strs)
	for i, v := range tempStr {
		r := []rune(v)
		slices.Sort(r)
		tempStr[i] = string(r)
	}

	res := [][]string{}
	seen := make(map[string]int)
	for i, v := range tempStr {
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
