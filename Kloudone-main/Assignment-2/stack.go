package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Stack struct {
	head *Node
	size int
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		fmt.Println("StackEmptyError")
		return 0
	}
	return s.head.data
}

func (s *Stack) Push(value int) {
	s.head = &Node{value, s.head}
	s.size++
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		fmt.Println("Stack is Empty")
		return
	}
	s.head = s.head.next
	s.size--
}

func (s *Stack) Print() {
	temp := s.head
	for temp != nil {
		fmt.Print(temp.data, " ")
		temp = temp.next
	}
	fmt.Println()
}

func main() {
	stk := new(Stack)
	stk.Push(20)
	stk.Push(35)
	stk.Push(22)
	stk.Push(33)
	stk.Print()
	stk.Pop()
	stk.Print()
}
