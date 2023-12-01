# ミューテックス

Go言語では、ゴルーチン（ `goroutines` ）を用いて並行処理を実装し、`sync.Mutex` を使用して排他制御（ `mutual exclusion` ）を行うことができます。以下のサンプルプログラムでは、複数のゴルーチンが共有リソースへのアクセスを競合する状況で、`sync.Mutex` を使用して排他処理を行います。

このプログラムでは、共有データ（この場合は単純なカウンタ）に対して複数のゴルーチンからのアクセスを同期させることで、データの一貫性を保持します。

## サンプルプログラム
```go
package main

import (
    "fmt"
    "sync"
    "time"
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
```

## 実行結果
```
Final counter value: 5000
```

このプログラムでは、`Counter` 型に `sync.Mutex` を組み込んでおり、`Increment` メソッドがカウンタを安全にインクリメントできるようにしています。メイン関数では、5つのゴルーチンを起動し、それぞれが1000回ずつカウンタをインクリメントします。`sync.WaitGroup` は、すべてのゴルーチンが完了するのを待つために使用されます。

このように `sync.Mutex` を使用することで、複数のゴルーチンが同じリソースへのアクセスを競合する際に、一貫性とデータの整合性を保持することができます。
