package main

import "fmt"

type NodeLL struct {
	data int
	Next *NodeLL
}

type StackLinkedList struct {
	head *NodeLL
	size int
}

func NewStackLinkedList() *StackLinkedList {
	return &StackLinkedList{}
}

func (s *StackLinkedList) IsEmpty() bool {
	return s.size == 0
}

func (s *StackLinkedList) Size() int {
	return s.size
}

func (s *StackLinkedList) Push(value int) {
	newNode := &NodeLL{data: value, Next: s.head}
	s.head = newNode
	s.size++
}

func (s *StackLinkedList) Pop() (int, bool) {
	if s.IsEmpty() {
		fmt.Println("StackLinkedList is empty!")
		return 0, false
	}

	data := s.head.data
	s.head = s.head.Next
	s.size--
	return data, true
}

func (s *StackLinkedList) Print() {
	tmp := s.head
	for tmp.Next != nil {
		fmt.Print(tmp.data, " ")
		tmp = tmp.Next
	}
	fmt.Println()
}

func main() {
	s := NewStackLinkedList()
	// Add value in stack
	s.Push(1)
	s.Push(21)
	s.Push(13)
	s.Push(11)
	s.Push(51)

	// print stack
	s.Print()

	// pop value of stack
	data, _ := s.Pop()
	println("Value ", data)

	data2, _ := s.Pop()
	println("Value ", data2)

	s.Print()
}
