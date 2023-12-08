# ラムダ式

Go言語では、ラムダ式（無名関数）を使用して、コードの一部で一時的な関数を定義し利用することができます。これは、高階関数の引数として関数を渡す場合や、即時関数として使用する場合に特に便利です。以下に、ラムダ式を使用したサンプルプログラムを示します。

## サンプルプログラム
**高階関数にラムダ式を渡す例**
```go
package main

import "fmt"

// Map はスライスの各要素に関数を適用します。
func Map[T any, U any](s []T, fn func(T) U) []U {
    var result []U
    for _, v := range s {
        result = append(result, fn(v))
    }
    return result
}

func main() {
    nums := []int{1, 2, 3, 4, 5}

    // Map 関数にラムダ式（無名関数）を渡して、各要素を2倍にする
    doubled := Map(nums, func(n int) int {
        return n * 2
    })

    fmt.Println(doubled) // [2 4 6 8 10]
}
```
このプログラムでは、`Map` 関数にラムダ式を渡しています。ラムダ式は、スライスの各要素を2倍にする処理を定義しています。

**即時関数の例**
```go
package main

import "fmt"

func main() {
    // 即時関数（IIFE）を定義し実行
    sum := func(nums ...int) int {
        total := 0
        for _, n := range nums {
            total += n
        }
        return total
    }(1, 2, 3, 4, 5)

    fmt.Println(sum) // 15
}
```

この例では、即時実行されるラムダ式（IIFE）を使用しています。ラムダ式は複数の整数を受け取り、それらの合計を計算して返します。

ラムダ式は、Go言語で動的な関数を定義するための強力なツールです。これにより、コードをよりコンパクトにし、特定のコンテキストでのみ使用される小さな関数を作成することができます。

