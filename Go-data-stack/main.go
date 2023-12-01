package main

import (
	"fmt"
	"sync"
)

// Stack はスタックを表します。
type Stack struct {
	items []interface{}
	mu    sync.Mutex
}

// NewStack は新しいスタックを作成します。
func NewStack() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
	}
}

// Push はスタックに要素を追加します。
func (s *Stack) Push(item interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, item)
}

// Pop はスタックから要素を取り出します。
func (s *Stack) Pop() interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.items) == 0 {
		return nil
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item
}

// IsEmpty はスタックが空かどうかを返します。
func (s *Stack) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.items) == 0
}

func main() {
	// スタックの作成。
	stack := NewStack()

	// スタックに要素を追加。
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// スタックから要素を取り出し、表示。
	fmt.Println(stack.Pop()) // 3
	fmt.Println(stack.Pop()) // 2
	fmt.Println(stack.Pop()) // 1
}
