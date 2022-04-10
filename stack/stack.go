package main

import "fmt"

type Stack struct {
	container []int
	top       int
	size      int
}

func NewStack(size int) *Stack {
	return &Stack{
		container: make([]int, size),
		top:       -1,
		size:      size,
	}
}

func (s *Stack) Push(val int) {
	if s.top == s.size-1 {
		fmt.Println("Stack is full")
	}
	s.top++
	s.container[s.top] = val
}

func (s *Stack) Pop() int {
	if s.top == -1 {
		fmt.Println("Stack is empty")
		return 0
	}
	val := s.container[s.top]
	s.top--
	return val
}

func main() {
	stack := NewStack(3)

	fmt.Println(stack.Pop())

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

}
