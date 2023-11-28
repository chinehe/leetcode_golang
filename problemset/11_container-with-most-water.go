package problemset

func maxArea(height []int) int {
	// 开始计算
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	maxArea := 0
	leftIndex := 0
	rightIndex := len(height) - 1
	for leftIndex < rightIndex {
		maxArea = max(maxArea, min(height[leftIndex], height[rightIndex])*(rightIndex-leftIndex))
		if height[leftIndex] > height[rightIndex] {
			rightIndex--
		} else {
			leftIndex++
		}
	}
	return maxArea
}

func maxArea2(height []int) int {
	type IV struct {
		i int
		v int
	}
	leftIVs := make([]IV, 0)
	rightIVs := make([]IV, 0)
	// 可能的左侧数据
	maxLeft := 0
	for i, v := range height {
		if v > maxLeft {
			maxLeft = v
			leftIVs = append(leftIVs, IV{i, v})
		}
	}
	// 可能的右侧数据
	maxRight := 0
	for i := len(height) - 1; i > 0; i-- {
		if height[i] > maxRight {
			maxRight = height[i]
			rightIVs = append(rightIVs, IV{i, height[i]})
		}
	}
	// 开始计算
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	maxArea := 0
	for _, leftIV := range leftIVs {
		for _, rightIV := range rightIVs {
			if leftIV.i >= rightIV.i {
				break
			}
			maxArea = max(maxArea, min(leftIV.v, rightIV.v)*(rightIV.i-leftIV.i))
		}
	}
	return maxArea
}
