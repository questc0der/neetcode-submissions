import "maps"
func isAnagram(s string, t string) bool {
	countS := make(map[string]int)
	countT := make(map[string]int)
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		countS[string(s[i])] = strings.Count(s, string(s[i]))
		countT[string(t[i])] = strings.Count(t, string(t[i]))
	}

	areEqual := maps.Equal(countS, countT)

	return areEqual
}
