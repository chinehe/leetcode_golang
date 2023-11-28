package problemset

func romanToInt(s string) int {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	num := 0
	preV := 0
	for _, r := range s {
		num += romanMap[r]
		if romanMap[r] > preV {
			num -= 2 * preV
		}
		preV = romanMap[r]
	}
	return num
}
