package main

import "fmt"

// creating node
type Node struct {
	data int
	next *Node
}

// creating list
type List struct {
	head  *Node
	count int
}

// function to return size
func (l *List) Size() int {
	return l.count
}

// function to check whether list is empty or not
func (l *List) IsEmpty() bool {
	return (l.count == 0)
}

// function to add element in the list
func (l *List) pushHead(value int) {
	l.head = &Node{value, l.head}
	l.count++
}

func (l *List) PrintList() {
	temp := l.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func (l *List) Lastpush(value int) {
	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}
}
func main() {
	mylist := &List{}
	mylist.Headpush(100)
	mylist.Headpush(200)
	mylist.Headpush(300)
	mylist.Headpush(155)
	mylist.Headpush(256)
	mylist.HeadPrintList()

}
