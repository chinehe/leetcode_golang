package problemset

func removeElement2(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}

func removeElement(nums []int, val int) int {
	endIndex := len(nums) - 1
	startIndex := 0
	for startIndex < endIndex {
		if nums[startIndex] == val {
			nums[startIndex] = nums[endIndex]
			endIndex--
		} else {
			startIndex++
		}
	}
	if startIndex == endIndex && nums[endIndex] == val {
		return endIndex
	}
	return endIndex + 1
}
