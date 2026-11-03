type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var result string
	for _, str := range strs {
		// Example: 5 + "#" + "hello" = "5#hello"
		result += fmt.Sprintf("%d#%s", len(str), str)
	}
	return result
}

func (s *Solution) Decode(encoded string) []string {
	var res []string
	i := 0
	
	for i < len(encoded) {
		// 1. Find where the delimiter '#' is
		j := i
		for encoded[j] != '#' {
			j++
		}
		
		// 2. Parse the length of the upcoming string
		length := 0
		fmt.Sscanf(encoded[i:j], "%d", &length)
		
		// 3. Extract the string using the length
		start := j + 1
		end := start + length
		res = append(res, encoded[start:end])
		
		// 4. Move the pointer past the extracted string
		i = end
	}
	
	return res
}
