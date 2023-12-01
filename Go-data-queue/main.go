package main

import (
	"fmt"
	"sync"
)

// Queue はキューを表します。
type Queue struct {
	items []interface{}
	mu    sync.Mutex
}

// NewQueue は新しいキューを作成します。
func NewQueue() *Queue {
	return &Queue{
		items: make([]interface{}, 0),
	}
}

// Enqueue はキューに要素を追加します。
func (q *Queue) Enqueue(item interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
}

// Dequeue はキューから要素を取り出します。
func (q *Queue) Dequeue() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.items) == 0 {
		return nil
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// IsEmpty はキューが空かどうかを返します。
func (q *Queue) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.items) == 0
}

func main() {
	// キューの作成。
	queue := NewQueue()

	// キューに要素を追加。
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	// キューから要素を取り出し、表示。
	fmt.Println(queue.Dequeue()) // 1
	fmt.Println(queue.Dequeue()) // 2
	fmt.Println(queue.Dequeue()) // 3
}
