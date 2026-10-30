func SortString(s string) string {
	// Convert string to a slice of runes
	r := []rune(s)
	// Sort the runes
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	// Convert back to string
	return string(r)
}

func groupAnagrams(strs []string) [][]string {
	sortedString := make(map[string][]string)
	for k, _ := range strs{
sortedString[SortString(strs[k])]=append(sortedString[SortString(strs[k])], strs[k])

	}
	var sublist [][]string

	for _, v := range sortedString {
		sublist = append(sublist, v)
	}
	return sublist
}
