package offer

// 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

//示例:
//
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL

//限制：
//
//0 <= 节点个数 <= 5000

func reverseListWithCopy(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newList := &ListNode{Val: head.Val}
	tmp := newList
	for head.Next != nil {
		head = head.Next
		tmp = &ListNode{Val: head.Val, Next: newList}
		newList = tmp
	}
	return newList
}

// 100.00% 100.00%
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead *ListNode
	for head != nil {
		tmp := head
		head = head.Next
		tmp.Next = newHead
		newHead = tmp
	}
	return newHead
}

// reverseListWithRecursion 递归
// 38.95% 12.14%
func reverseListWithRecursion(head *ListNode) *ListNode {
	return recursion(head, nil)
}

func recursion(head *ListNode, pre *ListNode) *ListNode {
	if head == nil {
		return pre
	}
	tmp := recursion(head.Next, head)
	head.Next = pre
	return tmp
}
