package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value interface{}
	next  *Node
	prev  *Node // add prev
}

type List struct {
	len  int
	head *Node
	tail *Node
}

type DoubleLinkedList interface {
	Init()
	Len() int
	PushFront(v interface{})
	PushBack(v interface{})
	Insert(v interface{}, index int) error
	RemoveFront() interface{}
	RemoveBack() interface{}
	Remove(index int) interface{}
	GetHead() *Node
	GetTail() *Node // add
	Get(index int) interface{}
}

func (l *List) Init() {
	l.len = 0
	l.head = nil
	l.tail = nil
}

func (l *List) Len() int {
	return l.len
}

func (l *List) PushFront(v interface{}) {
	node := &Node{
		value: v,
		next:  nil,
		prev:  nil, // add
	}
	if l.len == 0 {
		l.tail = node
	} else {
		l.head.prev = node // add
		node.next = l.head
	}
	l.head = node
	l.len++
}

func (l *List) PushBack(v interface{}) {
	node := &Node{
		value: v,
		next:  nil,
		prev:  nil, // add
	}
	if l.len == 0 {
		l.head = node
	} else {
		node.prev = l.tail // add
		l.tail.next = node
	}
	l.tail = node
	l.len++
}

func (l *List) Insert(v interface{}, index int) error {
	// new
	if index >= l.len {
		return errors.New("index error")
	}
	if index == 0 {
		l.PushFront(v)
		return nil
	}
	if index == l.len-1 {
		l.PushBack(v)
		return nil
	}

	mid := l.len / 2 // chose to start from the front or behind

	newNode := &Node{
		value: v,
		next:  nil,
		prev:  nil,
	}
	node := &Node{}
	if index <= mid { // from front
		node = l.head
		for i := 1; i <= mid; i++ {
			if i == index {
				break
			}
			node = node.next
		}
	} else { // from behind
		node = l.tail
		for i := l.len; i > mid; i-- {
			if i == index {
				break
			}
			node = node.prev
		}
	}
	newNode.prev = node
	newNode.next = node.next
	node.next.prev = newNode
	node.next = newNode
	l.len++
	return nil
}

func (l *List) RemoveFront() interface{} {
	if l.len == 0 {
		return nil
	}
	node := l.head
	l.head = l.head.next
	l.head.prev = nil // add
	l.len--
	return node.value
}

func (l *List) RemoveBack() interface{} {
	// new
	if l.len == 0 {
		return nil
	}
	node := l.tail
	l.tail = l.tail.prev
	l.tail.next = nil
	l.len--
	return node.value
}

func (l *List) Remove(index int) interface{} {
	// new
	if l.len == 0 {
		return nil
	}
	if index == 0 {
		return l.RemoveFront()
	} else if index == l.len-1 {
		return l.RemoveBack()
	}

	mid := l.len / 2

	node := &Node{}
	if index <= mid {
		node = l.head
		for i := 1; i <= mid; i++ {
			if i == index {
				break
			}
			node = node.next
		}

	} else {
		node = l.tail
		for i := l.len; i > mid; i-- {
			if i == index {
				break
			}
			node = node.prev
		}
	}
	tmp := node.next
	node.next = tmp.next
	tmp.next.prev = node
	l.len--
	return tmp.value
}

func (l *List) GetHead() *Node {
	return l.head
}

// add
func (l *List) GetTail() *Node {
	return l.tail
}

func (l *List) Get(index int) interface{} {
	// new
	if index >= l.len {
		return nil
	}

	mid := l.len

	node := &Node{}
	if index <= mid {
		node = l.head
		for i := 0; i <= mid; i++ {
			if i == index {
				return node.value
			}
			node = node.next
		}
	} else {
		node = l.tail
		for i := l.len; i > mid; i-- {
			if i == index {
				return node.value
			}
			node = node.prev
		}
	}
	return nil
}

func main() {
	var dl DoubleLinkedList = &List{}
	dl.Init()
	dl.PushFront(1)
	dl.PushFront(2)
	dl.PushFront(3)
	dl.PushBack(4)
	dl.PushBack(5)
	dl.PushBack(6)
	_ = dl.Insert(9, 1)
	_ = dl.Insert(9, 2)
	_ = dl.Insert(8, 5)
	_ = dl.Insert(8, 6)
	dl.RemoveFront()
	dl.RemoveBack()
	dl.Remove(1)
	dl.Remove(4)
	fmt.Println(1, dl.Get(1))
	fmt.Println(5, dl.Get(5))
	fmt.Println(3, dl.Get(3))

	fmt.Println("len", dl.Len())
	node := dl.GetHead()
	for {
		fmt.Print(node.value, " ")
		if node.next == nil {
			break
		}
		node = node.next
	}
	fmt.Print("\n")

	node = dl.GetTail()
	for {
		fmt.Print(node.value, " ")
		if node.prev == nil {
			break
		}
		node = node.prev
	}
	fmt.Print("\n")

}
