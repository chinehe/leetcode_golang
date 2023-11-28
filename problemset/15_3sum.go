package problemset

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	cur := 1
	for i := 0; i <= len(nums)-3; i++ {
		if nums[i] > 0 {
			break
		}
		if cur == nums[i] {
			continue
		}
		cur = nums[i]

		start := i + 1
		end := len(nums) - 1
		for end > start {
			sum := nums[start] + nums[end] + nums[i]
			switch {
			case sum == 0:
				res = append(res, []int{cur, nums[start], nums[end]})
				end--
				tmp := nums[start]
				for start = start + 1; start < len(nums)-1; start++ {
					if nums[start] != tmp {
						break
					}
				}
			case sum > 0:
				end--
			case sum < 0:
				tmp := nums[start]
				for start = start + 1; start < len(nums)-1; start++ {
					if nums[start] != tmp {
						break
					}
				}
			}
		}

	}
	return res
}
