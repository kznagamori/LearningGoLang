package main

import (
	"fmt"
	"sync"
)

// Counter は共有されるリソース（カウンタ）を表します。
type Counter struct {
	sync.Mutex
	Value int
}

// Increment はカウンタをインクリメントするメソッドです。
// Mutexを使用して排他制御を行います。
func (c *Counter) Increment() {
	c.Lock()
	c.Value++
	c.Unlock()
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	// 5つのゴルーチンを起動し、それぞれがカウンタをインクリメント
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	// すべてのゴルーチンが終了するのを待つ
	wg.Wait()

	// 最終的なカウンタの値を表示
	fmt.Println("Final counter value:", counter.Value)
}
