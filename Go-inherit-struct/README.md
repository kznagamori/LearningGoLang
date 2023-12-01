# 構造体の継承を実現する

Go言語にはクラスの継承という概念はありませんが、構造体の埋め込み（ `embedding` ）を利用して継承のような機能を実現することができます。これにより、ある構造体が別の構造体のフィールドやメソッドを「継承」するような振る舞いを実現できます。

以下に、構造体の埋め込みを使用した継承のサンプルプログラムを示します。

## サンプルプログラム
この例では、`Base` という基本構造体を作成し、それを `Derived` という派生構造体に埋め込みます。

**`Base` 構造体とそのメソッド**
```go
package main

import "fmt"

// Base は基本構造体です
type Base struct {
    Name string
}

// Greet は Base 構造体のメソッドです
func (b Base) Greet() string {
    return "こんにちは、" + b.Name + "です！"
}
```

**`Derived` 構造体（ `Base` を埋め込み）**
```go
// Derived は Base 構造体を埋め込んだ派生構造体です
type Derived struct {
    Base    // 匿名フィールドとして Base 構造体を埋め込む
    Degree string
}

// Introduce は Derived 構造体のメソッドです
func (d Derived) Introduce() string {
    return d.Greet() + " 私は" + d.Degree + "の学生です。"
}
```

**メイン関数**
```go
func main() {
    // Derived 構造体のインスタンスを作成
    d := Derived{
        Base: Base{Name: "Taro"},
        Degree: "コンピュータサイエンス",
    }

    // Base の Greet と Derived の Introduce メソッドを呼び出す
    fmt.Println(d.Greet())        // Base のメソッド
    fmt.Println(d.Introduce())    // Derived のメソッド
}
```

このプログラムでは、`Base` 構造体の `Greet` メソッドが `Derived` 構造体によって「継承」されています。`Derived` 構造体のインスタンスは、`Base` のメソッドに加えて、`Derived` 独自の `Introduce` メソッドも呼び出すことができます。

Goにおける構造体の埋め込みは、伝統的なオブジェクト指向言語の継承とは異なりますが、類似の振る舞いを実現するための強力な手段です。また、Goではインターフェースを使ってポリモーフィズムを実現することも一般的です。



