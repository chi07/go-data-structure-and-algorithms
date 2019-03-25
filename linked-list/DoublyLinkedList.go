package main

import "fmt"

type NodeInt struct {
	data int
	next *NodeInt
	prev *NodeInt
}

type DoublyLinkedList struct {
	head  *NodeInt
	tail  *NodeInt
	count int
}

func NewDoublyLinkListInt() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (dl *DoublyLinkedList) IsEmpty() bool {
	return dl.count == 0
}

func (dl *DoublyLinkedList) Size() int {
	return dl.count
}

func (dl *DoublyLinkedList) AddHead(value int) {
	newNode := &NodeInt{value, nil, nil}
	if dl.IsEmpty() {
		dl.head = newNode
		dl.tail = newNode
	} else {
		dl.head.prev = newNode
		newNode.next = dl.head
		dl.head = newNode
	}
}

func (dl *DoublyLinkedList) AddTail(value int) {
	newNode := &NodeInt{value, nil, nil}
	if dl.IsEmpty() {
		dl.head = newNode
		dl.tail = newNode
	} else {
		newNode.prev = dl.tail
		dl.tail.next = newNode
		dl.tail = newNode
	}
	dl.count--
}

func (dl *DoublyLinkedList) RemoveHead() (int, bool) {
	if dl.IsEmpty() {
		fmt.Println("List is empty!")
		return 0, false
	}

	value := dl.head.data
	dl.head = dl.head.next

	if dl.head == nil {
		dl.tail = nil
	} else {
		dl.head.prev = nil
	}

	return value, true
}

func (dl *DoublyLinkedList) RemoveNode(data int) bool {
	current := dl.head
	if current == nil {
		fmt.Println("Linked List is empty!")
		return false
	}

	// if head contain data which matched with key
	if current.data == data {
		current = current.next
		dl.count--
		if current != nil {
			dl.head = current
			dl.head.prev = nil
		} else { // there is only 1 node
			dl.tail = nil
		}
		return true
	}

	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			if current.next == nil { // last node
				dl.tail = current
			} else {
				current.next.prev = current
			}
			dl.count--
		}
		current = current.next
	}

	return false
}

func (dl *DoublyLinkedList) IsPresent(value int) bool {
	tmp := dl.head
	for tmp != nil {
		if tmp.data == value {
			return true
		}
		tmp = tmp.next
	}
	return false
}

func (dl *DoublyLinkedList) Print() {
	tmp := dl.head
	for tmp != nil {
		print(tmp.data, " ")
		tmp = tmp.next
	}
	fmt.Println()
}

func main() {
	l := NewDoublyLinkListInt()
	l.AddHead(11)
	l.AddHead(1)
	l.AddTail(22)
	l.AddTail(33)

	l.Print()

	l.RemoveHead()
	l.Print()

	l.RemoveNode(22)
	l.Print()

	fmt.Println(l.IsPresent(33))
}
