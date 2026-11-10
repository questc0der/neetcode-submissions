func isPalindrome(s string) bool {
	var reversedString strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		reversedString.WriteByte(s[i])
	}
	finalReversedString := reversedString.String()

	cleaned := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			return r 
		}
		return -1 
	}, s)

	revcleaned := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			return r 
		}
		return -1 
	}, finalReversedString)

	if strings.ToLower(cleaned) == strings.ToLower(revcleaned) {
		return true
	}
	return false
}
