package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedListInt struct {
	head  *Node
	count int
}

func NewLinkedListInt() *LinkedListInt {
	return &LinkedListInt{}
}

func (l *LinkedListInt) Size() int {
	return l.count
}

func (l *LinkedListInt) IsEmpty() bool {
	return l.count == 0
}

// Insert a node at node at the beginning of Singly Linked List
func (l *LinkedListInt) AddHead(value int) {
	l.head = &Node{data: value, next: l.head}
	l.count++
}

// Remove a node from beginning of linked list
func (l *LinkedListInt) RemoveHead() (int, bool) {
	if l.IsEmpty() {
		fmt.Println("Linked List is empty!")
		return 0, false
	}

	data := l.head.data
	l.head = l.head.next
	l.count--
	return data, true
}

// Insert a node at node at last of Singly Linked List
func (l *LinkedListInt) AddTail(data int) {
	newNode := &Node{data: data, next: nil}
	current := l.head

	if current == nil {
		l.head = newNode
	}

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (l *LinkedListInt) PrintLinkedList() {
	tmp := l.head
	for tmp != nil {
		fmt.Print(tmp.data, " ")
		tmp = tmp.next
	}
	fmt.Println("")
}

func (l *LinkedListInt) DeleteNode(value int) bool {
	tmp := l.head
	if l.IsEmpty() {
		fmt.Println("List is empty!")
		return false
	}

	// check if delete value is found of head
	if value == l.head.data {
		l.head = l.head.next
		l.count--
		return true
	}

	for tmp.next != nil {
		if tmp.next.data == value {
			tmp.next = tmp.next.next
			l.count--
			return true
		}
	}

	return false
}

func (l *LinkedListInt) IsPresent(value int) bool {
	tmp := l.head
	for tmp.next != nil {
		if tmp.next.data == value {
			return true
		}
		tmp = tmp.next
	}
	return false
}

func main() {
	l := NewLinkedListInt()
	l.AddHead(11)
	l.AddHead(1)
	l.AddTail(22)
	l.AddTail(33)

	l.PrintLinkedList()

	l.RemoveHead()
	l.PrintLinkedList()

	l.DeleteNode(22)
	l.PrintLinkedList()

	fmt.Println(l.IsPresent(33))
}
