# 構造体のインタフェースを使用したポリモーフィズムの実現

Go言語ではインターフェースを使用してポリモーフィズムを実現します。インターフェースは一連のメソッドを定義し、異なる型がこれらのメソッドを実装することで、同じインターフェースを共有できます。以下に、インターフェースを用いたポリモーフィズムの実現例を示します。

## サンプルプログラム
この例では、`Greet` という単一のメソッドを持つ `Greeter` インターフェースを定義し、異なる構造体がこのインターフェースを実装します。

**`Greeter` インターフェース**
```go
package main

import "fmt"

// Greeter は挨拶をするメソッドを持つインターフェースです。
type Greeter interface {
    Greet() string
}
```

**いくつかの構造体と `Greeter` インターフェースの実装**
```go
// Person 構造体
type Person struct {
    Name string
}

// Person 型に対する Greeter インターフェースの実装
func (p Person) Greet() string {
    return "こんにちは、私の名前は" + p.Name + "です！"
}

// Dog 構造体
type Dog struct {
    Name string
}

// Dog 型に対する Greeter インターフェースの実装
func (d Dog) Greet() string {
    return d.Name + "は「ワンワン！」と吠えます。"
}
```

**インターフェースを使用してポリモーフィズムを実現**
```go
func main() {
    // Greeter インターフェースを実装する異なる型のスライス
    greeters := []Greeter{
        Person{Name: "Taro"},
        Dog{Name: "ポチ"},
    }

    // 各 Greeter の Greet メソッドを呼び出す
    for _, greeter := range greeters {
        fmt.Println(greeter.Greet())
    }
}
```

このプログラムでは、`Person` と `Dog` の両方が `Greeter` インターフェースを実装しています。これにより、異なる型（ `Person` と `Dog` ）が同じインターフェース（ `Greeter` ）に基づいて同じ方法で扱われることが可能になります。`main` 関数では、`Greeter` インターフェースを実装するオブジェクトのスライスを作成し、それらの `Greet` メソッドを呼び出しています。

Goにおいて、インターフェースを使用することで、型の詳細を隠蔽し、異なる型が同じインターフェースに従って操作されるようにすることができます。これは、Goにおけるポリモーフィズムの強力な実現方法です。











