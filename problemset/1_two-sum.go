package problemset

func twoSum(nums []int, target int) []int {
	valueIndexMap := make(map[int]int)
	for index, value := range nums {
		if i, ok := valueIndexMap[value]; ok {
			return []int{i, index}
		}
		valueIndexMap[target-value] = index
	}
	return []int{}
}
