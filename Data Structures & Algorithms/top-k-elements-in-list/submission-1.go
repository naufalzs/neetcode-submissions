import "slices"

type pair struct {
	num int
	count int
}

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	
	pairs := []pair{}
	for k, v := range freq {
		pairs = append(pairs, pair{k, v})
	}

	slices.SortFunc(pairs, func(a,b pair) int {
		return b.count - a.count
	})

	res:= make([]int, 0, k)
	for i:= 0; i < k; i++ {
		res = append(res, pairs[i].num)
	}
	return res
}
