# defer
Go言語の `defer` ステートメントは、関数の実行が終了するとき（関数が `return` する時、もしくは `panic` が発生した時）に実行される関数やステートメントを指定するために使用されます。`defer` は特にリソースの解放やクリーンアップ作業に便利です。`defer` で指定された関数やステートメントは、それらが宣言された逆順で実行されます。

以下に、`defer` を使用したサンプルプログラムを示します。このプログラムでは、関数の開始時と終了時にメッセージを出力するシンプルな例を示しています。これにより、`defer` がいつ実行されるかがはっきりとわかります。

## サンプルプログラム
```go
package main

import "fmt"

func main() {
    fmt.Println("main function started")
    defer fmt.Println("main function deferred call")

    fmt.Println("main function running")

    exampleFunction()
    fmt.Println("main function ending")
}

func exampleFunction() {
    fmt.Println("exampleFunction started")
    defer fmt.Println("exampleFunction deferred call")
    fmt.Println("exampleFunction running")
}
```

このプログラムを実行すると、以下の順序でメッセージが出力されます：

1. `main function started`
2. `main function running`
3. `exampleFunction started`
4. `exampleFunction running`
5. `exampleFunction deferred call` (この時点でexampleFunctionが終了)
6. `main function ending`
7. `main function deferred call` (この時点でmain関数が終了)

この出力から、`defer` で指定されたステートメントが、それぞれの関数の終了時に実行されていることがわかります。`exampleFunction` の `defer` が先に実行され、その後 `main` 関数の `defer` が実行されます。これは `defer` が **LIFO（Last In, First Out）**の順序で実行されるためです。また、`defer`は関数の終了時に実行されるため、通常のステートメントよりも後に実行されます。

