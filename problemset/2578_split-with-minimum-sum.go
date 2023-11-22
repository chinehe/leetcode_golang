package problemset

import (
	"sort"
)

func splitNum(num int) int {
	nums := make([]int, 0)
	// 提取数字中的所有数字
	var mod int
	for num > 0 {
		mod = num % 10
		num = num / 10
		if mod == 0 {
			continue
		}
		nums = append(nums, mod)
	}
	// 对数字进行排序
	sort.Ints(nums)
	// 计算最小和
	var num1 int
	var num2 int
	halfLen := len(nums) / 2
	for i := 0; i < halfLen; i++ {
		num1 = num1*10 + nums[i*2]
		num2 = num2*10 + nums[i*2+1]
	}
	if len(nums)%2 == 1 {
		num1 = num1*10 + nums[len(nums)-1]
	}
	return num1 + num2
}
