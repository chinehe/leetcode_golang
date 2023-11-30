package problemset

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 遍历一次求取长度
	var size int
	tmp := head
	for tmp != nil {
		size++
		tmp = tmp.Next
	}
	// 单节点链表删除后必为空
	if size == 1 {
		return nil
	}
	// 计算要删除的位置
	deleteIndex := size - n
	// 删除首节点
	if deleteIndex == 0 {
		return head.Next
	}
	// 删除其他节点
	tmp = head
	// 将临时指针移动到要删除节点的上一个节点
	for deleteIndex > 1 {
		tmp = tmp.Next
		deleteIndex--
	}
	tmp.Next = tmp.Next.Next
	return head
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	list := make([]*ListNode, 0)
	for head != nil {
		list = append(list, head)
		head = head.Next
	}
	// 单节点链表删除后必为空
	if len(list) == 1 {
		return nil
	}
	// 根据要删除节点的位置来处理
	deleteIndex := len(list) - n
	switch deleteIndex {
	case 0:
		return list[1]
	case len(list) - 1:
		list[len(list)-2].Next = nil
		return list[0]
	default:
		list[deleteIndex-1].Next = list[deleteIndex].Next
		return list[0]
	}
}
