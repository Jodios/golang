package stack

import (
	"container/list"
	"fmt"
)

type Stack[T any] struct {
	stack *list.List
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		stack: list.New(),
	}
}

func (c *Stack[T]) Push(value T) {
	c.stack.PushFront(value)
}

func (c *Stack[T]) Pop() error {
	if c.stack.Len() > 0 {
		ele := c.stack.Front()
		c.stack.Remove(ele)
	}
	return fmt.Errorf("Pop Error: Stack is empty")
}

func (c *Stack[T]) Front() (T, error) {
	var r T
	if c.stack.Len() > 0 {
		if val, ok := c.stack.Front().Value.(T); ok {
			return val, nil
		}
		return r, fmt.Errorf("peep error: stack datatype is incorrect")
	}
	return r, fmt.Errorf("peep error: stack is empty")
}

func (c *Stack[T]) Size() int {
	return c.stack.Len()
}

func (c *Stack[T]) Empty() bool {
	return c.stack.Len() == 0
}
