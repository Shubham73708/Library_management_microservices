package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Peek() int {
	if q.IsEmpty() {
		fmt.Println("Queue is Empty")
		return 0
	}
	return q.head.data
}

func (q *Queue) Print() {
	temp := q.head
	for temp != nil {
		fmt.Print(temp.data, " ")
		temp = temp.next
	}
	fmt.Println()
}

func (q *Queue) Add(value int) {
	temp := &Node{value, nil}
	if q.head == nil {
		q.head = temp
		q.tail = temp
	} else {
		q.tail.next = temp
		q.tail = temp
	}
	q.size++
}

func (q *Queue) Remove() {
	if q.IsEmpty() {
		fmt.Println("QueueEmptyError")
		return
	}
	q.head = q.head.next
	q.size--
}

func main() {
	queue := new(Queue)
	queue.Add(100)
	queue.Add(450)
	queue.Add(190)
	queue.Add(149)
	queue.Add(354)
	queue.Print()
	queue.Remove()
	queue.Print()

}
