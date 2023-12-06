# リンクドリスト

Go言語でリンクドリストを使用するためのサンプルプログラムを作成します。この例では、標準ライブラリのcontainer/listパッケージを利用して、リンクドリストを操作する基本的な関数をいくつか実装します。具体的には、リストの作成、要素の追加、要素の削除、リストの表示などの操作を行います。

## サンプルプログラム
```go
package main

import (
    "container/list"
    "fmt"
)

func main() {
    // リンクドリストの作成
    l := list.New()

    // 要素の追加
    l.PushBack(1) // 末尾に追加
    l.PushFront(2) // 先頭に追加
    l.PushBack(3) // 末尾に追加

    // リンクリストの表示
    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }

    // 特定の要素の削除
    for e := l.Front(); e != nil; e = e.Next() {
        if e.Value == 2 {
            l.Remove(e)
            break
        }
    }

    fmt.Println("After removing an element:")

    // 要素削除後のリストの表示
    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }
}
```

## 実行結果
```
2
1
3
After removing an element:
1
3
```
このプログラムは、Go言語の標準ライブラリである `container/list` を使用しています。リンクドリストに対して、要素の追加、走査、削除などの操作を行い、その結果をコンソールに表示します。プログラムを実行すると、リストの初期状態と、特定の要素を削除した後の状態が出力されます。
