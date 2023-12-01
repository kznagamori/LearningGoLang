package main

import (
	"fmt"
	"sync"
	"time"
)

// doWork は独立したタスクを表します。
func doWork(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 処理が完了したらWaitGroupに通知

	fmt.Printf("Goroutine %d is starting\n", id)
	time.Sleep(time.Second) // 何かの処理を模擬
	fmt.Printf("Goroutine %d is done\n", id)
}

func main() {
	var wg sync.WaitGroup

	// 3つのゴルーチンを起動
	for i := 1; i <= 3; i++ {
		wg.Add(1) // WaitGroupのカウンタを増やす
		go doWork(i, &wg)
	}

	// すべてのゴルーチンが終了するのを待つ
	wg.Wait()
	fmt.Println("All goroutines complete.")
}
