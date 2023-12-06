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

// Get は指定されたインデックスの要素を参照します。
func (rb *RingBuffer) Get(index int) interface{} {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	if index < 0 || index >= rb.size {
		return nil
	}

	return rb.buffer[(rb.start+index)%rb.size]
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

	// バッファから要素を取り出し。
	fmt.Println("Pop:", rb.Pop()) // 1
	fmt.Println("Pop:", rb.Pop()) // 2

	// バッファに要素を追加。
	rb.Push(4)
	rb.Push(5)

	// バッファの要素を参照。
	fmt.Println("Get(0):", rb.Get(0)) // 3
	fmt.Println("Get(1):", rb.Get(1)) // 4
}
