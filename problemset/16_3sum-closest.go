package problemset

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	size := len(nums)
	sort.Ints(nums)

	minGap := math.MaxInt
	var bestSum int
	var gap int

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	minGapCal := func(sum int) {
		gap = abs(sum - target)
		if gap < minGap {
			minGap = gap
			bestSum = sum
		}
	}
	moveRight := func(cur, end int) int {
		next := cur + 1
		for ; next < end; next++ {
			if nums[next] != nums[cur] {
				break
			}
		}
		return next
	}
	moveLeft := func(cur, end int) int {
		next := cur - 1
		for ; next > end; next-- {
			if nums[next] != nums[cur] {
				break
			}
		}
		return next
	}

	// 三个指针
	var firstIndex int
	var secondIndex int
	var thirdIndex int
	var sum int
	for firstIndex < size-2 {
		secondIndex = firstIndex + 1
		thirdIndex = size - 1
		for thirdIndex > secondIndex {
			sum = nums[firstIndex] + nums[secondIndex] + nums[thirdIndex]
			switch {
			case sum == target:
				return sum
			case sum > target:
				minGapCal(sum)
				thirdIndex = moveLeft(thirdIndex, secondIndex)
			case sum < target:
				minGapCal(sum)
				secondIndex = moveRight(secondIndex, thirdIndex)
			}
		}
		firstIndex = moveRight(firstIndex, size-2)
	}
	return bestSum
}
