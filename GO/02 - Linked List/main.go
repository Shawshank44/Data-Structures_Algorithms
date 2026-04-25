package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Insert(value int) {
	newNode := &Node{data: value}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	temp := ll.head
	for temp.next != nil {
		temp = temp.next
	}

	temp.next = newNode
}

func (ll *LinkedList) Search(value int) bool {
	temp := ll.head
	for temp != nil {
		if temp.data == value {
			return true
		}
		temp = temp.next
	}
	return false
}

func (ll *LinkedList) Print() {
	temp := ll.head
	for temp != nil {
		fmt.Print(temp.data, " -> ")
		temp = temp.next
	}
	fmt.Println("nil")
}

func (ll *LinkedList) Update(oldValue, newValue int) bool {
	temp := ll.head

	for temp != nil {
		if temp.data == oldValue {
			temp.data = newValue
			return true
		}

		temp = temp.next
	}
	return false
}

func (ll *LinkedList) Delete(value int) bool {
	if ll.head == nil {
		return false
	}
	// Case 1 : head needs to be deleted
	if ll.head.data == value {
		ll.head = ll.head.next
		return true
	}

	// Case 2 : delete non-head node
	prev := ll.head
	curr := ll.head.next
	for curr != nil {
		if curr.data == value {
			prev.next = curr.next
			return true
		}
		prev = curr
		curr = curr.next
	}
	return false
}

func main() {
	ll := LinkedList{}

	// Create :
	ll.Insert(10)
	ll.Insert(20)
	ll.Insert(30)
	ll.Insert(40)
	ll.Insert(50)
	ll.Insert(99)
	ll.Print()

	// Read :
	fmt.Println("Search 20 : ", ll.Search(20))
	fmt.Println("Search 99 : ", ll.Search(99))

	// Update:
	ll.Update(20, 25)
	ll.Print()

	// Delete:
	ll.Delete(10)
	ll.Delete(30)
	ll.Print()
}
