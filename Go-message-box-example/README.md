# メッセージボックス

Go言語には、メッセージボックスのようなデータ共有のための組み込みパッケージはありませんが、チャネル（ `channels` ）とゴルーチン（ `goroutines` ）を使って、タイムアウト付きのメッセージボックスの機能を実装することができます。

以下に、メッセージボックスを模擬し、タイムアウト機能を持つ実行可能なサンプルプログラムを示します。この例では、メインゴルーチンがワーカーゴルーチンにメッセージを送信し、ワーカーゴルーチンがそれを処理するシナリオを実装しています。タイムアウトは `time.After` を使用して実装します。

## サンプルプログラム
```go
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
```

## 実行結果
```
Received message: Hello, World!
Message was processed successfully
```

このプログラムでは、`worker` 関数はメッセージボックス（ `messageBox` チャネル）を監視し、メッセージが到着するか、タイムアウト（3秒後）するまで待ちます。メッセージが到着すると、それを処理し、`done` チャネルに成功の信号を送信します。タイムアウトの場合は、失敗の信号を送信します。

メイン関数では、ワーカーゴルーチンを起動し、メッセージボックスにメッセージを送信して、ワーカーゴルーチンの処理結果を待ちます。これにより、メッセージボックスとタイムアウトの機能を模擬しています。
