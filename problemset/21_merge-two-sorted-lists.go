package problemset

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 特殊情况判定
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	// 首节点判断,让list1的首节点小于list2的首节点
	if list1.Val > list2.Val {
		tmp := list1
		list1 = list2
		list2 = tmp
	}
	// 往list1里面插入list2的值
	tmp := list1
	for tmp.Next != nil && list2 != nil {
		if tmp.Next.Val <= list2.Val {
			tmp = tmp.Next
		} else {
			t := tmp.Next
			tmp.Next = list2
			list2 = list2.Next
			tmp = tmp.Next
			tmp.Next = t
		}
	}
	if list2 != nil {
		tmp.Next = list2
	}
	return list1
}
