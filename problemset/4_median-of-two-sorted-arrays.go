package problemset

// 1 3 5 7
// 2 4 6 8
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	leftIndex := (length - 1) / 2
	rightIndex := length / 2
	return float64((findValueWithIndex(nums1, nums2, leftIndex) + findValueWithIndex(nums1, nums2, rightIndex)) / 2)
}

func findValueWithIndex(nums1 []int, nums2 []int, index int) int {
	//index1 := index / 2
	//index2 := index - index1
	//
	//for {
	//	if {
	//
	//	}
	//}
	return 0
}
