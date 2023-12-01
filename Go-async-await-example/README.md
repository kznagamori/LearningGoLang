# Async/Await

Go言語にはC#のような `async/await` 機能は直接存在しませんが、ゴルーチン（ `goroutines` ）とチャネル（ `channels` ）を使用して、非同期処理の流れを実装することができます。以下に、非同期処理を模擬するサンプルプログラムを示します。

このプログラムでは、ゴルーチンを起動して非同期的なタスクを実行し、タスクの結果をチャネルを通じて受け取ります。これは `async/await` のパターンに似ています。

## サンプルプログラム
```go
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
```

## 実行結果
```
Task 2 started
Task 1 started
Task 1 finished
Task 2 finished
Result of Task 1
Result of Task 2
```

このプログラムでは、`asyncTask` 関数がゴルーチンを起動し、非同期タスクを実行します。各タスクは2秒後に結果をチャネルに送信し、そのチャネルを返します。メイン関数では、これらのチャネルから結果を受け取ることで、非同期タスクの完了を待ちます。これはC#の `await` に相当します。

Go言語のゴルーチンとチャネルは、非同期処理を扱う強力なツールです。それらを使用することで、複雑な非同期処理や並行処理のパターンを実装することができます。
