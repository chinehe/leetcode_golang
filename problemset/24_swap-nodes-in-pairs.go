package problemset

func swapPairs(head *ListNode) *ListNode {
	// 节点数为0或1，直接返回
	if head == nil || head.Next == nil {
		return head
	}
	res := new(ListNode)
	res.Next = head
	tmp := res
	for head != nil && head.Next != nil {
		// 保存后面的元素
		t := head.Next.Next
		// tmp 指向第二个元素
		tmp.Next = head.Next
		// 第二个元素指向第一个元素
		tmp.Next.Next = head
		// 第一个元素指向后面的元素
		head.Next = t
		// 指针后移两位
		tmp = tmp.Next.Next
		head = tmp.Next
	}
	return res.Next
}
