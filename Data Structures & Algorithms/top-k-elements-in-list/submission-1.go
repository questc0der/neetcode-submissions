import (
	"cmp"
	"slices"
)

func topKFrequent(nums []int, k int) []int {
	numsMap := make(map[int]int)

	for _, num := range nums {
		numsMap[num]++
	}

	type pair struct {
		Key   int
		Value int
	}


	pairs := make([]pair, 0, len(numsMap))
	for key, val := range numsMap {
		pairs = append(pairs, pair{key, val})
	}

	slices.SortFunc(pairs, func(a, b pair) int {
		return cmp.Compare(b.Value, a.Value)
	})

	numm := make([]int, 0, k)
	for i := 0; i < k; i++ {
		numm = append(numm, pairs[i].Key)
	}

	return numm
}
