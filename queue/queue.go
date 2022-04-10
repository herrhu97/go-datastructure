package main

import "fmt"

type Queue struct {
	container []int

	front int

	rear int

	size int
}

func NewQueue(size int) *Queue {
	return &Queue{
		container: make([]int, size),
		front:     0,
		rear:      0,
		size:      0,
	}
}

func (q *Queue) IsEmpty() bool {
	if q.size == 0 {
		return true
	}
	return false
}

func (q *Queue) IsFull() bool {
	if q.size == len(q.container) {
		return true
	}
	return false
}

func (q *Queue) EnQueue(a int) bool {
	if q.IsFull() {
		return false
	}
	q.container[q.rear%len(q.container)] = a
	q.rear = q.rear%len(q.container) + 1
	q.size++
	return true
}

func (q *Queue) DeQueue() (bool, int) {
	if q.IsEmpty() {
		return false, 0
	}

	ret := q.container[q.front%len(q.container)]
	q.front = q.front%len(q.container) + 1
	q.size--
	return true, ret
}

func main() {
	queue := NewQueue(10)

	for j := 0; j < 3; j++ {
		for i := 0; i < 5; i++ {
			queue.EnQueue(i)
		}

		for i := 0; i < 6; i++ {
			fmt.Println(queue.DeQueue())
		}
	}

}
