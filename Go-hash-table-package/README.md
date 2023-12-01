# ハッシュテーブル

Go言語では、ハッシュテーブルの機能を提供する組み込みのデータ構造として `map` があります。`map` はキーと値のペアを保持し、高速な検索、挿入、削除操作を提供します。以下に、Goで `map` を使用してハッシュテーブルの基本的な操作を行うサンプルプログラムを示します。

## サンプルプログラム
```go
package main

import (
    "fmt"
)

func main() {
    // ハッシュテーブル（map）の初期化
    hashTable := make(map[string]int)

    // 要素の挿入
    hashTable["apple"] = 5
    hashTable["banana"] = 10
    hashTable["orange"] = 15

    // 要素のアクセスと表示
    fmt.Println("Apple count:", hashTable["apple"])
    fmt.Println("Banana count:", hashTable["banana"])

    // 要素の存在チェック
    value, exists := hashTable["grape"]
    if exists {
        fmt.Println("Grape count:", value)
    } else {
        fmt.Println("Grape not found")
    }

    // 要素の削除
    delete(hashTable, "banana")

    // ハッシュテーブルの内容を表示
    fmt.Println("Current hash table:", hashTable)
}
```

## 実行結果
```
Apple count: 5
Banana count: 10
Grape not found
Current hash table: map[apple:5 orange:15]
```
このプログラムでは、文字列をキーとし、整数を値とするハッシュテーブル（ `map` ）を作成しています。それから、いくつかの要素を挿入し、キーを使って値をアクセスしています。また、特定のキーがハッシュテーブルに存在するかどうかをチェックし、最後にキーを指定して要素を削除しています。

Goの `map` は非常に効率的で使いやすいハッシュテーブルの実装を提供し、キーと値のペアを簡単に管理できます。このサンプルプログラムは、Goでハッシュテーブルを使用する基本的な方法を示しています。
