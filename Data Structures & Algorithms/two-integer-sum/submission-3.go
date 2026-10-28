func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)
	for i := 0; i < len(nums);i++ {
		needed := target - nums[i]
		index, exists := numMap[needed]
		if exists && i != index{
			return []int{index, i}
		}else {
			numMap[nums[i]] = i
		}
	}
	return []int{}
}
