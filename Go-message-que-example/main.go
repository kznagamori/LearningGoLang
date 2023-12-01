package main

import (
	"fmt"
	"time"
)

func producer(queue chan<- string, totalMessages int) {
	for i := 1; i <= totalMessages; i++ {
		message := fmt.Sprintf("message%d", i)
	loop:
		for {
			select {
			case queue <- message:
				fmt.Printf("Produced: %s\n", message)
				break loop
			case <-time.After(100 * time.Millisecond): // キューが満杯の場合のタイムアウト
				fmt.Println("Queue is full, producer timeout")
			}
			time.Sleep(100 * time.Millisecond) // 次のメッセージを送るまでの遅延
		}

	}
}

func consumer(queue <-chan string) {
	for {
		select {
		case message := <-queue:
			fmt.Printf("Consumed: %s\n", message)
			time.Sleep(1 * time.Second) // 消費の遅延をシミュレート
		case <-time.After(2 * time.Second):
			fmt.Println("No messages, consumer timeout")
			return
		}
	}
}

func main() {
	queueSize := 3
	queue := make(chan string, queueSize)
	totalMessages := 10

	go producer(queue, totalMessages)
	go consumer(queue)

	time.Sleep(15 * time.Second) // メインゴルーチンが終了するのを防ぐ
}
