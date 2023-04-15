package linkedList

import "fmt"

type Node struct {
	prev  *Node
	next  *Node
	value interface{} // TODO: make this typed
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func New() LinkedList {
	return LinkedList{}
}

func (l *LinkedList) Insert(value interface{}) {
	list := &Node{
		prev:  l.tail,
		value: value,
	}

	// if tail exists set its next value to new node
	if l.tail != nil {
		l.tail.next = list
	} else {
		list.next = nil
		l.head = list
	}

	l.tail = list
}

func (l *LinkedList) Delete(value interface{}) {
	list := l.head

	for list != nil {
		if list.value == value {
			list.prev.next = list.next
			list.next.prev = list.prev

		}
		list = list.next
	}
}

// TODO Implement find and refactor delete to call find
func (l *LinkedList) Find(value interface{}) {

}

func (l *LinkedList) Print() {
	list := l.head
	fmt.Print("Head ->")
	for list != nil {
		fmt.Printf("%v -> ", list.value)
		list = list.next
	}
	fmt.Println(" Tail")
}
