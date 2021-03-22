package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
}

func (list *List) add(value int) {
	newNode := &Node{value, nil}
	if list.head == nil {
		list.head = newNode
		fmt.Println("Node added sussuccfully")
		return
	}
	curr := list.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
	fmt.Println("Node added sussuccfully")
}

func (list *List) addAtHead(data int) {
	newNode := &Node{data, nil}
	if list.head == nil {
		list.head = newNode
		fmt.Println("Node added at head sussuccfully")
		return
	}
	newNode.next = list.head
	list.head = newNode
	fmt.Println("Node added at head sussuccfully")
}

func (list *List) printList() {
	curr := list.head
	for curr != nil {
		fmt.Print(curr.data, " ")
		curr = curr.next
	}
	fmt.Println()
}

func (list *List) delete(data int) {
	if list.head == nil {
		fmt.Println("list is Empty")
		return
	}
	curr := list.head
	if list.head.data == data {
		list.head = list.head.next
		fmt.Println("Node is deleted")
		return
	}
	for curr.next != nil {
		if curr.next.data == data {
			curr.next = curr.next.next
			fmt.Println("Node is deleted")
			return
		}
		curr = curr.next
	}
	fmt.Println("Node is not found")
}

func main() {
	list := new(List)
	list.add(100)
	list.add(200)
	list.add(300)
	list.add(500)
	list.printList()
	list.addAtHead(2)
	list.printList()
	list.delete(29)
	list.printList()
}
