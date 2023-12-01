# スレッド

Go言語では、スレッドの代わりにゴルーチン（ `goroutines` ）を使用して並行処理を行います。ゴルーチンは軽量なスレッドのようなもので、Goランタイムによって管理されます。以下に、ゴルーチンを使用して並行処理を行うサンプルプログラムを示します。

このプログラムでは、複数のゴルーチンを起動して、それぞれが独立して処理を実行します。また、`sync.WaitGroup` を使用して、メイン関数がすべてのゴルーチンの完了を待つことができるようにしています。


## サンプルプログラム
```go
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
```

## 実行結果
```
Goroutine 3 is starting
Goroutine 1 is starting
Goroutine 2 is starting
Goroutine 2 is done
Goroutine 1 is done
Goroutine 3 is done
All goroutines complete.
```

このプログラムでは、`main` 関数内で `sync.WaitGroup` のインスタンスを作成し、3つのゴルーチンを起動しています。各ゴルーチンは `doWork` 関数を実行し、処理が終わると `wg.Done()` を呼び出して、`WaitGroup` に処理の完了を通知します。メイン関数は `wg.Wait()` を呼び出して、すべてのゴルーチンが完了するのを待ちます。

このようにして、Go言語ではゴルーチンと `sync.WaitGroup` を使用して効率的に並行処理を実装することができます。ゴルーチンは非常に軽量で、多数のゴルーチンを同時に動作させることが可能です。
