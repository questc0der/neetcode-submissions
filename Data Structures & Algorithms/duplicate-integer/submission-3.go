func hasDuplicate(nums []int) bool {
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++{
		_, occurred := numMap[nums[i]]
		if occurred {
			return true
		}
		numMap[nums[i]] = nums[i]
	}
	return false
}
