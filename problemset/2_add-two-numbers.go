package problemset

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := new(ListNode)
	tmp := res
	v1 := 0
	v2 := 0
	s := 0
	p := 0
	for l1 != nil || l2 != nil || p != 0 {
		v1 = 0
		v2 = 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		s = v1 + v2 + p
		p = s / 10
		tmp.Next = &ListNode{
			Val: s % 10,
		}
		tmp = tmp.Next
	}
	return res.Next
}
