package offer

// 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
// 输入：head = [1,3,2]
// 输出：[2,3,1]
// 0 <= 链表长度 <= 10000

// ListNode 单向链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// reversePrintWithDoubleSlice 双Slice
// 4ms 3MB
func reversePrintWithDoubleSlice(head *ListNode) []int {
	values := make([]int, 0)
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	res := make([]int, len(values))
	for i := 0; i < len(values); i++ {
		res[len(values)-1-i] = values[i]
	}
	return res
}

// reversePrintWithCount 先计数后取值
// 4ms 2.6MB
func reversePrintWithCount(head *ListNode) []int {
	count := 0
	temp := head
	for temp != nil {
		count++
		temp = temp.Next
	}
	res := make([]int, count)
	index := 1
	for head != nil {
		res[count-index] = head.Val
		index++
		head = head.Next
	}
	return res
}

type Stack struct {
	values []int
	len    int
}

func (s *Stack) Push(val int) {
	if len(s.values) == s.len {
		s.values = append(s.values, val)
	} else {
		s.values[s.len] = val
	}
	s.len++
}

func (s *Stack) Pop() (val int) {
	s.len--
	return s.values[s.len]
}

// reversePrintWithStack use stack
// 0ms 3MB
func reversePrintWithStack(head *ListNode) []int {
	stack := &Stack{
		values: make([]int, 0),
		len:    0,
	}
	for head != nil {
		stack.Push(head.Val)
		head = head.Next
	}
	res := make([]int, stack.len)
	for i := range res {
		res[i] = stack.Pop()
	}
	return res
}

// reversePrintWithRecursion use recursion
// 4ms 4.4MB
func reversePrintWithRecursion(head *ListNode) []int {
	recursionRes := new([]int)
	recur(head, recursionRes)
	return *recursionRes
}

func recur(head *ListNode, res *[]int) {
	if head != nil {
		recur(head.Next, res)
		*res = append(*res, head.Val)
	}
}
