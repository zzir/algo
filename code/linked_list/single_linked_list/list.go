package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value interface{}
	next  *Node
}

type List struct {
	len  int
	head *Node
	tail *Node
}

type SingleLinkedList interface {
	Init()
	Len() int
	PushFront(v interface{})
	PushBack(v interface{})
	Insert(v interface{}, index int) error
	RemoveFront() interface{}
	RemoveBack() interface{}
	Remove(index int) interface{}
	GetHead() *Node
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
	}
	if l.len == 0 {
		l.tail = node
	} else {
		node.next = l.head
	}
	l.head = node
	l.len++
}

func (l *List) PushBack(v interface{}) {
	node := &Node{
		value: v,
		next:  nil,
	}
	if l.len == 0 {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
	l.len++
}

func (l *List) Insert(v interface{}, index int) error {
	if index >= l.len {
		return errors.New("index error")
	}
	if index == 0 {
		l.PushFront(v)
		return nil
	}
	node := &Node{
		value: v,
		next:  nil,
	}
	prev := l.head
	for i := 1; i < l.len; i++ {
		if i == index {
			node.next = prev.next
			prev.next = node
			l.len++
			return nil
		}
		prev = prev.next
	}

	return errors.New("insert error")
}

func (l *List) RemoveFront() interface{} {
	if l.len == 0 {
		return nil
	}
	node := l.head
	l.head = l.head.next
	l.len--
	return node.value
}

func (l *List) RemoveBack() interface{} {
	if l.len == 0 {
		return nil
	}
	node := l.head
	for i := 1; i < l.len-1; i++ {
		node = node.next
	}
	tmp := node.next
	node.next = nil
	l.tail = node
	l.len--
	return tmp.value
}

func (l *List) Remove(index int) interface{} {
	if l.len == 0 {
		return nil
	}
	if index == 0 {
		return l.RemoveFront()
	} else if index == l.len-1 {
		return l.RemoveBack()
	}
	node := l.head
	for i := 1; i < index; i++ {
		node = node.next
	}
	tmp := node.next
	node.next = node.next.next
	l.len--
	return tmp.value
}

func (l *List) GetHead() *Node {
	return l.head
}

func (l *List) Get(index int) interface{} {
	if index >= l.len {
		return nil
	}
	node := l.head
	for i := 0; i < l.len; i++ {
		if i == index {
			return node.value
		}
		node = node.next
	}
	return nil
}

func main() {
	var sl SingleLinkedList = &List{}
	sl.Init()
	sl.PushFront(6)
	sl.PushFront(7)
	sl.PushFront(8)
	sl.PushBack(1)
	sl.PushBack(2)
	sl.PushBack(3)
	fmt.Println("r", sl.RemoveFront())
	_ = sl.Insert(9, 0)
	_ = sl.Insert(8, 5)
	fmt.Println("r", sl.RemoveBack())
	fmt.Println("r", sl.RemoveBack())
	fmt.Println("rm", sl.Remove(1))
	fmt.Println("rm", sl.Remove(2))

	fmt.Println("len", sl.Len())
	node := sl.GetHead()
	for {
		fmt.Println(">>>", node.value)
		if node.next == nil {
			break
		}
		node = node.next
	}
	fmt.Println(sl.Get(1))
}
