package doubleLinkedList

import "fmt"

// 双向链表
// 带有前驱和后继结点的线性表
// 双向链表是带有头结点的

type node struct {
	prev, next *node
	data       interface{}
}

func (n *node) Prev() *node {
	return n.prev
}

func (n *node) Next() *node {
	return n.next
}

func (n *node) GetValue() interface{} {
	return n.data
}

type DoubleList struct {
	length int
	head   *node
	tail   *node
}

func New() *DoubleList {
	head := new(node)
	tail := head
	list := new(DoubleList)
	list.head = head
	list.tail = tail
	list.length = 0
	return list
}

func (d *DoubleList) Head() *node {
	return d.head
}

func (d *DoubleList) Tail() *node {
	return d.tail
}

func (d *DoubleList) PrintForward() {
	for temp := d.Head().Next(); temp != d.Head(); temp = temp.Next() {
		fmt.Println(temp.GetValue())
	}
}

func (d *DoubleList) PrintBack() {
	for temp := d.Tail(); temp != d.Head(); temp = temp.Prev() {
		fmt.Println(temp.GetValue())
	}
}

// 在双向链表中, prev 和 next 均是指向节点,而非前后节点的 prev 或者 next
func (d *DoubleList) Append(data interface{}) {
	e := new(node)
	e.data = data
	if d.length == 0 {
		d.head.next = e
		d.head.prev = e
		e.prev = d.head
		e.next = d.head
		d.tail = e
		d.length++
	} else {
		d.tail.next = e
		d.head.prev = e
		e.prev = d.tail
		e.next = d.head
		d.tail = e
		d.length++
	}
}
