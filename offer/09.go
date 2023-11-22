package offer

// 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
// 分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead操作返回 -1 )

//输入：
//["CQueueWithLinkedList","appendTail","deleteHead","deleteHead","deleteHead"]
//[[],[3],[],[],[]]
//输出：[null,null,3,-1,-1]

//输入：
//["CQueueWithLinkedList","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
//[[],[],[5],[2],[],[]]
//输出：[null,-1,null,null,5,2]

// 1 <= values <= 10000
// 最多会对 appendTail、deleteHead 进行 10000 次调用

// CQueueWithLinkedList use linked list
// 74.18% 51.43%
type CQueueWithLinkedList struct {
	head *ListNode
	tail *ListNode
}

func CQueueWithLinkedListConstructor() CQueueWithLinkedList {
	return CQueueWithLinkedList{}
}

func (this *CQueueWithLinkedList) AppendTail(value int) {
	temp := &ListNode{Val: value}
	if this.head == nil {
		this.head = temp
		this.tail = temp
	} else {
		this.tail.Next = temp
		this.tail = temp
	}
}

func (this *CQueueWithLinkedList) DeleteHead() int {
	if this.head == nil {
		return -1
	}
	val := this.head.Val
	if this.head.Next == nil {
		this.head = nil
		this.tail = nil
	} else {
		this.head = this.head.Next
	}
	return val
}

type CQueueWithDoubleStack struct {
	head Stack
	tail Stack
}

// CQueueWithDoubleStackConstructor use double stack
// 53.67% 76.74%
func CQueueWithDoubleStackConstructor() CQueueWithDoubleStack {
	return CQueueWithDoubleStack{
		head: Stack{
			values: make([]int, 0),
		},
		tail: Stack{
			values: make([]int, 0),
		},
	}
}

func (this *CQueueWithDoubleStack) AppendTail(value int) {
	this.tail.Push(value)
}

func (this *CQueueWithDoubleStack) DeleteHead() int {
	if this.head.len > 0 {
		return this.head.Pop()
	}
	if this.tail.len > 0 {
		for this.tail.len > 0 {
			this.head.Push(this.tail.Pop())
		}
		return this.head.Pop()
	}
	return -1
}
