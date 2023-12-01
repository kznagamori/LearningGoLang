# リングバッファ

Go言語におけるリングバッファ（環状バッファ）を使用する例として、標準ライブラリの `container/ring` パッケージを利用したサンプルプログラムを示します。このプログラムでは、リングバッファの初期化、要素の追加、要素の走査、要素の操作などの基本的な機能を実装します。

## サンプルプログラム
```go
package main

import (
    "container/ring"
    "fmt"
)

func main() {
    // リングバッファのサイズを指定して初期化
    r := ring.New(5)

    // リングバッファにデータを追加
    for i := 0; i < r.Len(); i++ {
        r.Value = i
        r = r.Next()
    }

    // リングバッファの要素を表示
    r.Do(func(p interface{}) {
        fmt.Println(p)
    })

    // 特定の要素を操作
    r.Move(2) // 2つ先に移動
    r.Value = 100

    fmt.Println("After modification:")

    // 変更後のリングバッファの要素を表示
    r.Do(func(p interface{}) {
        fmt.Println(p)
    })
}

```

## 実行結果
```
0
1
2
3
4
After modification:
100
1
2
3
4
```
