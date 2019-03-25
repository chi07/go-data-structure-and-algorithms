package main

import (
	"errors"
	"fmt"
)

type NodeItem struct {
	data int
	Next *NodeItem
}

type QueueLinkedList struct {
	Head *NodeItem
	Tail *NodeItem
	size int
}

func NewQueueLinkedList() *QueueLinkedList {
	return &QueueLinkedList{}
}

func (q *QueueLinkedList) IsEmpty() bool {
	return q.size == 0
}

func (q *QueueLinkedList) Size() int {
	return q.size
}

func (q *QueueLinkedList) Enqueue(value int) {
	newNode := &NodeItem{data: value, Next: nil}
	if q.Head == nil {
		q.Head = newNode
		q.Tail = newNode
	} else {
		q.Tail.Next = newNode
		q.Tail = newNode
	}
	q.size++
}

func (q *QueueLinkedList) Dequeue() int {
	if q.IsEmpty() {
		return 0
	}

	data := q.Head.data
	q.Head = q.Head.Next
	q.size--
	return data
}

// Peek function returns the first item in the queue without removing it
func (q *QueueLinkedList) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("QueueLinkedList in empty")
	}
	return q.Head.data, nil
}

func (q *QueueLinkedList) Print() {
	curr := q.Head
	for curr != nil {
		fmt.Print(curr.data, " ")
		curr = curr.Next
	}
	println()
}

func main() {
	q := NewQueueLinkedList()
	// add items into queue
	for i := 1; i <= 20; i++ {
		q.Enqueue(i)
	}
	// print queue
	q.Print()
	println(q.Size())

	// get first item data of queue
	first, err := q.Peek()
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	println(first)

	// dequeue 2 first items
	q.Dequeue()
	q.Dequeue()

	q.Print()
	println(q.Size())
}
