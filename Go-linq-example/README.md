# C#みたいなLINQ機能

Go言語にはC#の `LINQ` のような組み込みの機能はありませんが、スライスに対して関数を適用することで `LINQ` のような操作を模倣することが可能です。これには、高階関数を使用します。以下に、フィルタ、マップ、リデュースの操作を行うサンプルプログラムを示します。

## サンプルプログラム
```go
package main

import (
    "fmt"
    "strings"
)

// Filter はスライスから条件に合致する要素だけを抽出します。
func Filter[T any](s []T, fn func(T) bool) []T {
    var result []T
    for _, v := range s {
        if fn(v) {
            result = append(result, v)
        }
    }
    return result
}

// Map はスライスの各要素に関数を適用します。
func Map[T any, U any](s []T, fn func(T) U) []U {
    var result []U
    for _, v := range s {
        result = append(result, fn(v))
    }
    return result
}

func main() {
    // スライスの初期化
    words := []string{"go", "rust", "java", "python"}

    // Filter: 文字列の長さが3以上の要素を抽出
    longWords := Filter(words, func(s string) bool {
        return len(s) > 3
    })
    fmt.Println("長さが3以上の単語:", longWords)

    // Map: 文字列を大文字に変換
    upperWords := Map(words, strings.ToUpper)
    fmt.Println("大文字に変換した単語:", upperWords)
}
```

このプログラムでは、`Filter` 関数と `Map` 関数を定義しています。`Filter` 関数は条件に合う要素のみを抽出し、Map関数は各要素に関数を適用します。これにより、LINQのような操作を模倣しています。

Go言語ではジェネリクスが `Go 1.18` で導入されたため、さまざまな型のスライスに対して同じ関数を使用することが可能になります。ただし、このサンプルプログラムは `Go 1.18` 以降でのみ動作します。`LINQ` のようなクエリ機能を完全に再現するのは難しいですが、このような関数を組み合わせることで、類似の操作を実現することができます。
