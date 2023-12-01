# 複数ファイルを使用する

複数のファイルを使用したプログラムを構成する際、一般的には機能ごとにファイルを分け、それぞれで異なるタスクを処理します。  
以下に、シンプルな例を示します。この例では、二つのファイルを使用します: `main.go` と `greetings.go`。`main.go` はプログラムのエントリポイントとなり、`greetings.go` は挨拶のロジックを含みます。

## ファイル構成
- main.go: メインプログラム。
- greetings.go: 挨拶を生成する関数を含む。

### greetings.go
このファイルには、挨拶を生成する関数を定義します。
```go
// greetings.go
package main

import "fmt"

// Greet は名前を受け取り、挨拶文を返します。
func Greet(name string) string {
    return fmt.Sprintf("こんにちは、%sさん!", name)
}
```

### main.go
このファイルはプログラムのエントリポイントです。`greetings.go`のGreet関数を使用します。

```go
// main.go
package main

import "fmt"

func main() {
    // Greet 関数を呼び出し
    greeting := Greet("Taro")
    fmt.Println(greeting)
}
```

## プログラムの実行
1. これらのファイルを同じディレクトリに配置します。
2. ディレクトリ内で `go run .` コマンドを実行します。これにより、`main.go` と `greetings.go` の両方がコンパイルされ、プログラムが実行されます。

この例では、Greet 関数は `greetings.go` ファイル内で定義され、`main.go` から呼び出されています。これにより、コードの構造が整理され、機能ごとにファイルを分割することができます。Goでは、同じパッケージに属するファイル間で関数や変数を自由に共有できます。
