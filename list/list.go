package main

import (
	"fmt"
	"strings"
)

type Node struct {
	E    int
	Next *Node
}

type List struct {
	dummyHead *Node
	size      int
}

func initNode(e int) *Node {
	return &Node{
		E:    e,
		Next: nil,
	}
}

func InitList() *List {
	return &List{
		dummyHead: initNode(0),
		size:      0,
	}
}

func (l *List) Head() *Node {
	return l.dummyHead
}

func (l *List) Size() int {
	return l.size
}

func (l *List) IsEmpty() bool {
	return l.size == 0
}

// 在链表的第index索引个元素后插入元素,索引从0开始
func (l *List) AddIndex(index, e int) {
	if index > l.size || index < 0 {
		panic("Add failed. Illegal index.")
	}

	prev := l.dummyHead
	node := initNode(e)

	for i := 0; i < index; i++ {
		prev = prev.Next
	}

	node.Next = prev.Next
	prev.Next = node
	l.size++
}

// 在链表头添加元素
func (l *List) AddFirst(e int) {
	l.AddIndex(0, e)
}

// 在链表尾部添加节点
func (l *List) AddLast(e int) {
	l.AddIndex(l.size, e)
}

// 在链表中查询是否包括元素e
func (l *List) Contains(e int) bool {
	cur := l.dummyHead.Next
	for cur != nil {
		if cur.E == e {
			return true
		}
		cur = cur.Next
	}
	return false
}

// 在链表中查询第index个元素
func (l *List) Get(index int) int {
	if index > l.size || index < 0 {
		panic("Get failed. Illegal index.")
	}

	cur := l.dummyHead.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.E
}

func (l *List) GetFirst() int {
	return l.Get(0)
}

func (l *List) GetLast() int {
	return l.Get(l.size - 1)
}

func (l *List) Reverse() {
	var pre *Node
	var cur *Node
	next := l.dummyHead.Next
	for next != nil { // 四行经典写法，翻转链表
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}
	l.Head().Next = pre
}

func (this *List) String() string {
	var builder strings.Builder
	cur := this.dummyHead.Next
	for cur != nil {
		fmt.Fprintf(&builder, "%d -> ", cur.E)
		cur = cur.Next
	}
	fmt.Fprintf(&builder, "NULL")
	return builder.String()
}

func main() {
	list := InitList()
	for i := 0; i < 5; i++ {
		list.AddFirst(i)
	}
	fmt.Println(list)
	list.Reverse()
	fmt.Println(list)
}
