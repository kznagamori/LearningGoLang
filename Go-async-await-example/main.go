package main

import (
	"fmt"
	"time"
)

// asyncTask は非同期に実行するタスクを表します。
// この関数は結果をチャネルに送信します。
func asyncTask(id int) <-chan string {
	resultChan := make(chan string)

	go func() {
		// タスクの実行をシミュレート
		fmt.Printf("Task %d started\n", id)
		time.Sleep(time.Second * 2) // 2秒待機
		fmt.Printf("Task %d finished\n", id)

		// 結果をチャネルに送信
		resultChan <- fmt.Sprintf("Result of Task %d", id)
		close(resultChan)
	}()

	return resultChan
}

func main() {
	// 非同期タスクの実行
	resultChan1 := asyncTask(1)
	resultChan2 := asyncTask(2)

	// タスクの結果を待つ（"await"相当）
	result1 := <-resultChan1
	result2 := <-resultChan2

	// 結果の表示
	fmt.Println(result1)
	fmt.Println(result2)
}
