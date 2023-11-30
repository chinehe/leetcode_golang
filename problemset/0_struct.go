package problemset

import (
	"errors"
)

// ListNode 单向链表
type ListNode struct {
	Val  int
	Next *ListNode
}

type stack struct {
	data []rune
	head int
	tail int
}

func newStack(size int) *stack {
	return &stack{
		data: make([]rune, size),
		head: 0,
		tail: 0,
	}
}

func (s *stack) push(v rune) {
	s.data[s.tail] = v
	s.tail++
}

func (s *stack) pop() (rune, error) {
	if s.tail == 0 {
		return 0, errors.New("stack is empty")
	}
	s.tail--
	return s.data[s.tail], nil
}

func (s *stack) size() int {
	return s.tail - s.head
}
