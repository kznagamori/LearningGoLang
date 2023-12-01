# ソート
Go言語では、`sort` パッケージを使用して様々な型のスライスをソートすることができます。以下のサンプルプログラムでは、整数や文字列のスライスをソートする方法と、カスタムの比較関数を使ってスライスをソートする方法を示します。

## サンプルプログラム
```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    // 整数のスライスをソート
    intSlice := []int{4, 2, 7, 3, 8}
    sort.Ints(intSlice)
    fmt.Println("Sorted integers:", intSlice)

    // 文字列のスライスをソート
    stringSlice := []string{"banana", "apple", "cherry"}
    sort.Strings(stringSlice)
    fmt.Println("Sorted strings:", stringSlice)

    // カスタム比較関数を使ってスライスをソート
    people := []struct {
        Name string
        Age  int
    }{
        {"Alice", 30},
        {"Bob", 25},
        {"Clara", 27},
    }

    // 年齢でソート
    sort.Slice(people, func(i, j int) bool {
        return people[i].Age < people[j].Age
    })

    fmt.Println("Sorted by age:", people)
}
```

## 実行結果
```
Sorted integers: [2 3 4 7 8]
Sorted strings: [apple banana cherry]
Sorted by age: [{Bob 25} {Clara 27} {Alice 30}]
```

このプログラムでは、まず標準の `sort` パッケージの `Ints` 関数と`Strings` 関数を使用して、整数と文字列のスライスをそれぞれソートしています。これらの関数はそれぞれ、整数のスライスと文字列のスライスに特化しており、使い方は非常にシンプルです。

次に、構造体のスライスをソートするために、`sort.Slice` 関数を使用しています。この関数は、スライスと比較関数を引数として受け取ります。比較関数は、スライス内の2つの要素を比較して、どちらが先に来るべきかを決定します。この例では、構造体のスライスを年齢でソートしています。

これらの方法を使用することで、Go言語でさまざまな型のスライスを簡単にソートすることができます。
