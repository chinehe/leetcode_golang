package problemset

import (
	"sort"
)

func nextPermutation(nums []int) {
	// 单个元素直接返回
	if len(nums) <= 1 {
		return
	}
	// 从倒数第二个元素开始倒序遍历
	for i := len(nums) - 2; i >= 0; i-- {
		// 从倒数第一个元素开始倒序遍历，直到当前 i 的下一个元素
		for j := len(nums) - 1; j > i; j-- {
			// 如果找到元素比 i 的元素大，就交换i,j位置，剩余的升序排列
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
				sort.Ints(nums[i+1:])
				return
			}
		}
	}
	// 走到这里说明数组是降序排列的，改成升序排列
	sort.Ints(nums)
}
