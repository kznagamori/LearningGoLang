# 外部パッケージを使用したプログラムを作成する手順

外部パッケージを使用したプログラムを作成する手順とサンプルプログラムを以下に示します。

**1. プロジェクトのセットアップ:**  

作業用ディレクトリを作成します（例： `mkdir Go-use-package` ）。  
ディレクトリに移動し（例： `cd Go-use-package` ）、 `go mod init [モジュール名]コマンド` （例： `go mod init Go-use-package` ）で新しいモジュールを初期化します。


**2. 外部パッケージのインストール:**

必要な外部パッケージを `go -u get [パッケージのパス]` コマンドでインストールします。例: `go get -u github.com/gorilla/mux`

**3. コードの記述:**

ソースファイル（例えばmain.go）を作成し、外部パッケージを含むプログラムを記述します。
```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    // ルーターの初期化
    r := mux.NewRouter()

    // エンドポイントとハンドラーの設定
    r.HandleFunc("/", HomeHandler)

    // サーバーの起動
    log.Fatal(http.ListenAndServe(":8000", r))
}

// ハンドラー関数
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, World!")
}

```
このプログラムでは、 `gorilla/mux` ルーターを使用してルートURL (/) にアクセスがあった場合に `Hello, World!` というメッセージを返す簡単なWebサーバーを作成します。  
このプログラムを実行した後、ブラウザで `http://localhost:8000/` にアクセスすると、メッセージが表示されます。

**4. ビルドと実行:**

`go build` コマンドでプログラムをビルドし、生成された実行ファイルを実行します。