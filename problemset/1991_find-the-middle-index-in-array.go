package problemset

func pivotIndex(nums []int) int {
	var rightSum int
	for i, _ := range nums {
		rightSum += nums[i]
	}
	var leftSum int

	for i, _ := range nums {
		rightSum -= nums[i]
		if leftSum == rightSum {
			return i
		}
		leftSum += nums[i]
	}

	return -1
}
