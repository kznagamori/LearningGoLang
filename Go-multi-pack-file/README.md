# 複数ファイルを機能ごとにパッケージを分けて使用するサンプルプログラム

複数のファイルにわたってパッケージを分けて使用するサンプルプログラムを作成する際、機能ごとに異なるパッケージを作成し、それらをメインアプリケーションでインポートして使用します。以下に、簡単な例を示します。この例では、二つのパッケージを使用します: `main` パッケージと `greetings` パッケージ。

## プロジェクト構成
- main.go: メインプログラム（main パッケージ）。
- greetings/greetings.go: 挨拶を生成する関数を含む（greetings パッケージ）。

## ディレクトリ構造
```
Go-multi-pack-file/
├── main.go
└── greetings/
    └── greetings.go
```

### greetings/greetings.go
`greetings` パッケージでは、挨拶を生成する関数を定義します。
```go
// greetings/greetings.go
package greetings

import "fmt"

// Greet は名前を受け取り、挨拶文を返します。
func Greet(name string) string {
    return fmt.Sprintf("こんにちは、%sさん!", name)
}
```

### main.go
`main` パッケージ（メインプログラム）では、`greetings` パッケージの `Greet` 関数を使用します。

```go
// main.go
package main

import (
    "fmt"
    "Go-multi-pack-file/greetings" // ローカルの `greetings` パッケージをインポート
)

func main() {
    // greetings パッケージの Greet 関数を呼び出し
    greeting := greetings.Greet("Taro")
    fmt.Println(greeting)
}
```

## プログラムの実行
上記のディレクトリ構造に従ってファイルを配置します。  
プロジェクトのルートディレクトリ（ `Go-multi-pack-file` ）で `go run . ` コマンドを実行します。  
この例では、greetings パッケージは `greetings/greetings.go` に定義されており、main パッケージ（ `main.go` ）からインポートして使用されています。  
Goでは、このように異なるパッケージを作成し、機能ごとにコードを整理することが一般的です。  
また、パッケージのインポートパスは、Goモジュールの名前から始まることに注意してください（この例では `Go-multi-pack-file/greetings` ）。





