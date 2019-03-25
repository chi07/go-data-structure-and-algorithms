package main

import "fmt"

type Node struct {
	data int
	Next *Node
}

type Stack struct {
	head *Node
	size int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Push(value int) {
	newNode := &Node{data: value, Next: s.head}
	s.head = newNode
	s.size++
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		fmt.Println("Stack is empty!")
		return 0, false
	}

	data := s.head.data
	s.head = s.head.Next
	s.size--
	return data, true
}

func (s *Stack) Print() {
	tmp := s.head
	for tmp.Next != nil {
		fmt.Print(tmp.data, " ")
		tmp = tmp.Next
	}
	fmt.Println()
}

func main() {
	s := NewStack()
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
