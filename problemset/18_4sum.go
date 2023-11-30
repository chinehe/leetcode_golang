package problemset

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	// 特殊值判定
	if len(nums) < 4 {
		return [][]int{}
	}
	// 排序
	sort.Ints(nums)
	// 左移右移函数
	moveR := func(cur, end int) int {
		for i := cur + 1; i < end; i++ {
			if nums[cur] != nums[i] {
				return i
			}
		}
		return end
	}
	moveL := func(cur, end int) int {
		for i := cur - 1; i > end; i-- {
			if nums[cur] != nums[i] {
				return i
			}
		}
		return end
	}
	// 双重循环+双指针
	res := make([][]int, 0)
	size := len(nums)
	var start int
	var end int
	var sum int
	for i := 0; i < size-3; {
		if nums[i]*4 > target {
			break
		}
		for j := i + 1; j < size-2; {
			if nums[i]+nums[j]*3 > target {
				break
			}
			start = j + 1
			end = size - 1
			for start < end {
				if nums[i]+nums[j]+nums[start]*2 > target {
					break
				}
				if nums[i]+nums[j]+nums[end]*2 < target {
					break
				}
				sum = nums[i] + nums[j] + nums[start] + nums[end]
				switch {
				case sum < target:
					start = moveR(start, end)
				case sum > target:
					end = moveL(end, start)
				default:
					res = append(res, []int{nums[i], nums[j], nums[start], nums[end]})
					start = moveR(start, end)
					end = moveL(end, start)
				}
			}
			j = moveR(j, size-2)
		}
		i = moveR(i, size-3)
	}
	return res
}
