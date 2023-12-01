package main

import (
	"fmt"
	"sync"
)

// RingBuffer はリングバッファを表します。
type RingBuffer struct {
	buffer []interface{}
	size   int
	start  int
	end    int
	mu     sync.Mutex
}

// NewRingBuffer は新しいリングバッファを作成します。
func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		buffer: make([]interface{}, size),
		size:   size,
	}
}

// Push はバッファに新しい要素を追加します。
func (rb *RingBuffer) Push(item interface{}) {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	rb.buffer[rb.end] = item
	rb.end = (rb.end + 1) % rb.size

	// バッファが満杯の場合、開始位置を移動します。
	if rb.end == rb.start {
		rb.start = (rb.start + 1) % rb.size
	}
}

// Pop はバッファから要素を取り出します。
func (rb *RingBuffer) Pop() interface{} {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	if rb.IsEmpty() {
		return nil
	}

	item := rb.buffer[rb.start]
	rb.start = (rb.start + 1) % rb.size
	return item
}

// IsEmpty はバッファが空かどうかを返します。
func (rb *RingBuffer) IsEmpty() bool {
	return rb.start == rb.end
}

func main() {
	// リングバッファの作成。
	rb := NewRingBuffer(5)

	// バッファに要素を追加。
	rb.Push(1)
	rb.Push(2)
	rb.Push(3)

	// バッファから要素を取り出し、表示。
	fmt.Println(rb.Pop()) // 1
	fmt.Println(rb.Pop()) // 2
	fmt.Println(rb.Pop()) // 3
}
