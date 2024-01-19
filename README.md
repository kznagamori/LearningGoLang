# Go言語学習

## Go言語とは
Go言語（または単にGoとも呼ばれる）は、Googleによって開発されたプログラミング言語です。2009年に公開され、効率的なコンパイル、実行時のパフォーマンス、および容易なプログラミングのためのシンプルなデザインが特徴です。以下にGo言語の主な特徴を挙げます：

1. コンパイル言語：Goはコンパイル言語であり、ソースコードは実行可能なバイナリファイルにコンパイルされます。これにより、高い実行速度が得られます。

2. 静的型付け：Goは静的型付け言語で、変数の型はコンパイル時に決定されます。これはプログラムの安全性を高め、エラーを早期に発見するのに役立ちます。

3. 並行性のサポート：Goは並行プログラミングを簡単にするための「ゴルーチン（goroutines）」と呼ばれる軽量スレッドの概念を導入しました。これにより、同時に多くのタスクを効率的に処理できます。

4. ガベージコレクション：Goには自動メモリ管理（ガベージコレクション）が含まれており、開発者はメモリの割り当てと解放を手動で行う必要がありません。

5. 標準ライブラリ：Goには豊富な標準ライブラリがあり、ネットワーキング、暗号化、データ操作など多くの共通タスクをカバーしています。

6. クロスプラットフォーム：Goは多くのプラットフォーム（Windows、Linux、macOSなど）で動作し、一度書いたコードをさまざまな環境で利用できます。

7. シンプルな構文：Goはシンプルで読みやすい構文を持っており、学習しやすいと考えられています。

これらの特性により、Goはシステムプログラミング、クラウドベースのアプリケーション、マイクロサービス、並行処理が必要なアプリケーションなど、多様な用途で使用されています。

### 推しポイント
- 構文がC言語体系とは異なるが、言語機能がシンプルなので、覚えやすい。
- 実行速度が速い。
- 使用可能なパッケージが多く存在するので、ツールを作成しやすい。
- ChatGPTで質問すると、かなり正解に近いコードを取得できる。

### いまいち、ダメなポイント
- GUIアプリ、スマフォアプリ作成には向いていない。

## 基礎
- [プロジェクト作成からビルドまでの手順](./Go-start-project/README.md)
- [外部パッケージを使用したプログラムを作成する手順](./Go-use-package/README.md)
- [複数ファイルを使用する](./Go-multi-file/README.md)
- [複数ファイルを機能ごとにパッケージを分けて使用する](./Go-multi-pack-file/README.md)
- [ファイル内グローバル変数と関数、ファイル外グローバル変数と関数](./Go-global-local-scope/README.md)
- [クラスを使う](./Go-struct-methods/README.md)
- [パブリックな構造体のメンバ、メソッドを定義とプライベートな構造体のメンバ、メソッドを定義する](./Go-public-private/README.md)
- [構造体の継承を実現する](./Go-inherit-struct/README.md)
- [構造体のインタフェースを使用したポリモーフィズムの実現](./Go-poly-structs/README.md)
- [C#みたいな構造体のプロパティを定義、使用](./Go-struct-property/README.md)
- [一般的なエラー処理](./Go-error-handling/README.md)
- [各型のstring formatの組み合わせ](./Go-string-format/README.md)
- [ジェネリック機能](./Go-generics-example/README.md)
- [ラムダ式](./Go-lambda-example/README.md)
- [C#みたいなLINQ機能](./Go-linq-example/README.md)


## アルゴリズムとデータ構造
- [リンクドリスト構造](./Go-linked-list/README.md)
- [リングバッファ構造](./Go-ring-buffer/README.md)
- [キュー構造](./Go-data-queue/README.md)
- [スタック構造](./Go-data-stack/README.md)
- [二分木構造](./Go-binary-tree/README.md)
- [平衡二分木構造](./Go-balanced-tree/README.md)
- [ハッシュテーブル](./Go-hash-table/README.md)
- [クイックソート](./Go-quick-sort/README.md)
- [再帰を使用しないシェルソート](./Go-non-recursive-shell-sort/README.md)

## 応用
- [リンクドリスト](./Go-list-package/README.md)
- [リングバッファ](./Go-ring-buffer-package/README.md)
- [ハッシュテーブル](./Go-hash-table-package/README.md)
- [ソート](./Go-sort-package/README.md)
- [スレッド](./Go-threading-example/README.md)
- [Async/Await](./Go-async-await-example/README.md)
- [排他処理](./Go-mutex-example/README.md)
- [メッセージボックス](./Go-message-box-example/README.md)
- [メッセージキュー](./Go-message-que-example/README.md)

## エコシステム
- [パッケージをインストールして使用する](./Go-install-package/README.md)

## 言語特性
- [defer](./Go-defer-example/README.md)

## サンプルプログラム

