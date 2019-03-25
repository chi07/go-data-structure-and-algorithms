package main

import (
	"fmt"
)

type Stack struct {
	Nodes []*Node
	Count uint
	Max   uint
}

type Node struct {
	Value int
}

func NewStack() *Stack {
	return &Stack{}
}

// Size function returns
func (s *Stack) Size() int {
	return len(s.Nodes)
}

// IsEmpty function return bool that check stack is empty or not
func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) IsFull() bool {
	return s.Count >= s.Max
}

// Push a node into stack
func (s *Stack) Push(n *Node) {
	if !s.IsFull() {
		s.Nodes = append(s.Nodes, n)
		s.Count++
		return
	}
	fmt.Println("Stack is full. Can not push into stack")
}

func (s *Stack) Pop() *Node {
	if s.IsEmpty() {
		fmt.Println("Stack is empty. Can not get any node")
		return nil
	}

	s.Count--
	return s.Nodes[s.Count]
}

// Top function returns last node of stack and not remove it from stack
func (s *Stack) Top() *Node {
	return s.Nodes[s.Size()-1]
}

func (s *Stack) PrintStack() {
	for i := 0; i < s.Size(); i++ {
		fmt.Print(s.Nodes[i].Value, " ")
	}
	fmt.Println()
}

func main() {
	s := NewStack()
	s.Max = 10

	exampleValues := [10]int{123, 13, 46672, 10, 35, 10, 163, 22, 177, 999}
	for _, v := range exampleValues {
		node := &Node{}
		node.Value = v
		s.Push(node)
	}

	s.PrintStack()

	fmt.Printf("Size of Stack: %d \n", s.Size())
	fmt.Printf("Get node in top of stack %v \n", s.Top().Value)

	fmt.Printf("Node get by Pop function: %v\n", s.Pop().Value)
	fmt.Printf("Node get by Pop function: %v\n", s.Pop().Value)
	fmt.Printf("Node get by Pop function: %v\n", s.Pop().Value)
	fmt.Printf("Node get by Pop function: %v\n", s.Pop().Value)
	fmt.Printf("Node get by Pop function: %v\n", s.Pop().Value)
	fmt.Printf("Node get by Pop function: %v\n", s.Pop().Value)
	fmt.Printf("Node get by Pop function: %v\n", s.Pop().Value)
	fmt.Printf("Node get by Pop function: %v\n", s.Pop().Value)
	fmt.Printf("Node get by Pop function: %v\n", s.Pop().Value)
	fmt.Printf("Node get by Pop function: %v\n", s.Pop().Value)
}
