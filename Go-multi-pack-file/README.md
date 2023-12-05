# 複数ファイルを機能ごとにパッケージを分けて使用する
Go言語で複数のファイルを機能ごとにパッケージに分けて使用し、ライブラリと呼び出し側のパッケージを別のプロジェクトとして構成する例を以下に示します。

## プロジェクト構成
### 1. ライブラリプロジェクト: mylibrary
**ディレクトリ構造**
```
mylibrary/
├── go.mod
└── greetings
    └── greetings.go
```
**mylibrary/go.mod**
```
module example.com/mylibrary

go 1.21.4
```
**mylibrary/greetings/greetings.go**
```go
package greetings

// Greet は挨拶を返します。
func Greet(name string) string {
    return fmt.Sprintf("こんにちは、%sさん!", name)
}
```

### 2. 呼び出し側のプロジェクト: myapp
**ディレクトリ構造**
```
myapp/
├── go.mod
└── main.go
```
**myapp/go.mod**
```
module example.com/myapp

go 1.21.4

require example.com/mylibrary v0.1.0

replace example.com/mylibrary => ../mylibrary
```
ここで、`example.com/mylibrary v0.1.0` は、`mylibrary` ライブラリの適切なバージョンに置き換えてください。ローカルで開発している場合、`replace` ディレクティブを使ってローカルパスを指定します。

**myapp/main.go**
```go
package main

import (
    "fmt"
    "example.com/mylibrary/greetings"
)

func main() {
    fmt.Println(greetings.Greet("Taro"))
}
```

## 手順
1. `mylibrary` と `myapp` プロジェクトをそれぞれ別のディレクトリに作成します。
2. `mylibrary` プロジェクトのルートで `go mod init example.com/mylibrary` を実行して `go.mod` ファイルを初期化します。
3. `myapp` プロジェクトのルートで `go mod init example.com/myapp` を実行し、`go.mod` ファイルを初期化します。
4. `myapp` の `go.mod` ファイルに `mylibrary` への依存関係を追加します。ローカル開発の場合は 'replace' ディレクティブを使用します。
5. `myapp` ディレクトリで `go run main.go` を実行して、ライブラリの関数が呼び出されることを確認します。

この構成により、`mylibrary` ライブラリを `myapp` プロジェクトで利用することができます。`go.mod` ファイルは、各プロジェクトの依存関係を管理するために使用され、ローカル開発中のパスの指定には `replace` ディレクティブが役立ちます。

