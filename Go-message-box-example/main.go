package main

import (
	"fmt"
	"time"
)

func worker(messageBox chan string, done chan bool) {
	for {
		select {
		case message := <-messageBox:
			// メッセージを受信
			fmt.Println("Received message:", message)
			done <- true
		case <-time.After(3 * time.Second):
			// タイムアウト
			fmt.Println("No message received: timeout")
			done <- false
			return
		}
	}
}

func main() {
	messageBox := make(chan string)
	done := make(chan bool)

	// ワーカーゴルーチンを起動
	go worker(messageBox, done)

	// メッセージボックスにメッセージを送信
	messageBox <- "Hello, World!"

	// ワーカーゴルーチンの終了を待つ
	if success := <-done; success {
		fmt.Println("Message was processed successfully")
	} else {
		fmt.Println("Failed to process message")
	}
}
