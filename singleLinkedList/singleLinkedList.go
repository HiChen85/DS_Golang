package linkedlist

import "fmt"

// 单链表节点
type node struct {
	value interface{}
	next  *node
}

func (n *node) Next() *node {
	return n.next
}

func (n *node) Value() interface{} {
	return n.value
}

type List struct {
	length int
	// 我们统一头结点不包含任何数据
	head *node // 头指针
	tail *node
}

func NewList() *List {
	l := new(List)
	l.head = nil
	l.tail = nil
	return l
}

func (l *List) Head() *node {
	return l.head
}

func (l *List) Len() int {
	return l.length
}

func (l *List) PrintValue() {
	for temp := l.head; temp != nil; temp = temp.next {
		fmt.Println(temp.value)
	}
}

func (l *List) Append(key interface{}) {
	e := new(node)
	e.value = key
	e.next = nil
	if l.head == nil {
		l.head = e
	} else {
		l.tail.next = e
	}
	l.tail = e
	l.length++
}

func (l *List) Insert(value, at interface{}) {
	e := new(node)
	e.value = value
	e.next = nil

	if l.head.value == at {
		e.next = l.head
		l.head = e
		l.length++
	} else {
		for prev, flag := l.head, l.head.next; flag != nil; {
			if flag.value != at {
				prev = prev.next
				flag = flag.next
			} else {
				prev.next = e
				e.next = flag
				l.length++
				return
			}
		}
		fmt.Println("插入位置有误")

	}
}

// 此函数只删除与 key 匹配的第一个数据
func (l *List) Delete(key interface{}) {

	if l.head.value == key {
		l.head = l.head.next
		l.length--
		//return nil
	} else {
		for prev, flag := l.head, l.head.next; flag != nil; {
			if flag.value != key {
				prev = prev.next
				flag = flag.next
			} else {
				prev.next = flag.next
				l.length--
				fmt.Printf("%v删除成功\n", key)
				return
				//return nil
			}
		}
	}
	fmt.Println("该数据不在链表中")

}
