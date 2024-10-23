package problemset

func countCompleteDayPairs(hours []int) int64 {
	counts := make([]int64, 24)
	// 取余
	for _, h := range hours {
		counts[h%24]++
	}
	total := int64(0)
	for i := 0; i <= 12; i++ {
		if i == 0 || i == 12 {
			total += (counts[i] * (counts[i] - 1)) / 2
		} else {
			total += counts[i] * counts[24-i]
		}
	}
	return total
}
