package main

import "fmt"

type Node struct {
	Value int
}

const MaxQueue = 10

type Queue struct {
	Nodes    []*Node
	Count    int
	Capacity int
	Front    int
	Back     int
}

func NewQueue() *Queue {
	return &Queue{Capacity: MaxQueue, Nodes: make([]*Node, MaxQueue)}
}

func (q *Queue) isEmpty() bool {
	return q.Count == 0
}

func (q *Queue) isFull() bool {
	return q.Count == q.Capacity
}

func (q *Queue) Size() int {
	return len(q.Nodes)
}

func (q *Queue) enqueue(n *Node) {
	if q.isFull() {
		fmt.Println("Queue is full. Can not add any node")
		return
	}
	q.Nodes[q.Back] = n
	q.Back = (q.Back + 1) % len(q.Nodes)
	q.Count++
}

func (q *Queue) dequeue() *Node {
	if q.isEmpty() {
		fmt.Println("Queue is empty!")
		return nil
	}

	q.Count--
	node := &Node{}
	node = q.Nodes[q.Front]
	q.Front = (q.Front + 1) % q.Size()

	return node
}

func main() {
	q := NewQueue()
	q.enqueue(&Node{Value: 12})
	q.enqueue(&Node{Value: 22})
	q.enqueue(&Node{Value: 33})
	q.enqueue(&Node{Value: 44})
	q.enqueue(&Node{Value: 55})
	q.enqueue(&Node{Value: 66})
	q.enqueue(&Node{Value: 77})
	q.enqueue(&Node{Value: 88})
	q.enqueue(&Node{Value: 99})
	q.enqueue(&Node{Value: 199})
	q.enqueue(&Node{Value: 299})

	fmt.Println("Queue: ", q)
	println()
	fmt.Println(q.dequeue(), q.dequeue(), q.dequeue(), q.dequeue(), q.dequeue(), q.dequeue(), q.dequeue(), q.dequeue())
	fmt.Println(q.dequeue(), q.dequeue(), q.dequeue(), q.dequeue())
}
