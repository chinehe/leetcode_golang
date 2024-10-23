package problemset

func removeDuplicates(nums []int) int {
	// 长度小于2，不需要删除
	if len(nums) < 2 {
		return len(nums)
	}
	// 找出所有重复的区间
	type indexRange struct {
		start int
		end   int
	}
	dupRanges := make([]indexRange, 0)
	start := 0
	end := 0
	for end < len(nums) {
		if nums[start] == nums[end] {
			end++
		} else {
			if end > start+1 {
				dupRanges = append(dupRanges, indexRange{start: start, end: end - 1})
			}
			start = end
			end++
		}
	}
	if end > start+1 {
		dupRanges = append(dupRanges, indexRange{start: start, end: end - 1})
	}
	// 根据重复区间删除元素
	deleteNums := 0
	for _, dupRange := range dupRanges {
		nums = append(nums[:dupRange.start-deleteNums+1], nums[dupRange.end-deleteNums+1:]...)
		deleteNums += dupRange.end - dupRange.start
	}
	return len(nums)
}
