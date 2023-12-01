# メッセージキュー
Go言語で指定された機能を持つメッセージキューを実装するために、チャネルとゴルーチンを使用します。このサンプルプログラムでは、キューの最大サイズを3とし、合計10個のデータを送信します。送信側はキューがいっぱいになるまでデータを送信し、キューが開くたびに新しいデータを送信します。受信側はキューから1秒に1回データを受信します。

## サンプルプログラム
```go
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

```

このGoプログラムは、メッセージキューの生産者（ `producer` ）と消費者（ `consumer` ）の間のデータのやり取りをシミュレートするものです。キューのサイズは3に設定されており、合計10個のメッセージが生産されます。プログラムの流れは次のようになります。

**プロデューサー (producer関数)**
- 合計10回のループを実行し、各ループで `messageX` という形式のメッセージを生成しキューに送信します。
- キューが満杯でメッセージが受け入れられない場合、100ミリ秒ごとにキューに空きが出るかチェックします。
- メッセージがキューに正常に送信された場合、`Produced: messageX` と出力し、ループを抜けます。
- メッセージ送信の試行中にキューが満杯の状態が続くと `Queue is full, producer timeout` と出力し、次の試行に移ります。

**コンシューマー (consumer関数)**
- メッセージキューからメッセージを受信し続けます。
- メッセージが受信されると、`Consumed: messageX` と出力し、1秒間スリープ（消費の遅延をシミュレート）します。
- 2秒間メッセージがキューに来ない場合、`No messages, consumer timeout` と出力し、関数から抜けます。

**メイン関数 (main関数)**
- バッファサイズ3の文字列チャネルを作成します。
- `producer`関数と`consumer`関数をそれぞれゴルーチンとして起動します。
- `time.Sleep(15 * time.Second)` により、メインゴルーチンが終了しプログラムが終了するのを15秒間待機します。これにより、`producer` と`consumer` ゴルーチンが十分に動作する時間を確保します。

このプログラムは、Go言語のチャネルとゴルーチンを使用した並行処理の基本的な例を示しており、メッセージキューの概念を理解するのに役立ちます。また、タイムアウト処理の実装方法も示しています。
