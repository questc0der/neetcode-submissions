func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)
	count := 0
	for count < len(nums){
		numMap[nums[count]] = count
		count++
	}
	for i := 0; i < len(nums);i++ {
		needed := target - nums[i]
		index, exists := numMap[needed]
		if exists && i != index{
			return []int{i, index}
		}
	}
	return []int{}
}
