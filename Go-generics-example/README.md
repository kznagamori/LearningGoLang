# ジェネリック機能

Go言語では、バージョン 1.18 からジェネリック機能が導入されました。ジェネリックを使用すると、型パラメータを用いて柔軟な関数や型を定義できます。以下に、ジェネリックを使用したサンプルプログラムを示します。

## ジェネリック関数の例

**generics_function.go**

```go
package main

import "fmt"

// Compare は2つの値を比較し、等しいかどうかを返します。
// この関数は任意の型（T）に対して動作します。
func Compare[T comparable](a, b T) bool {
    return a == b
}

func main() {
    fmt.Println(Compare[int](5, 5))   // true
    fmt.Println(Compare[string]("go", "rust")) // false
}
```

このプログラムでは、`Compare` 関数は任意の `comparable` 型（Goの基本型のいずれか）に対して動作します。この関数は2つの値を比較し、それらが等しいかどうかをブール値で返します。

## ジェネリック型の例

**generics_type.go**
```go
package main

import "fmt"

// Box は任意の型の値を保持する汎用コンテナです。
type Box[T any] struct {
	Value T
}

// Set はBoxに値を設定します。
func (b *Box[T]) Set(value T) {
	b.Value = value
}

// Get はBoxから値を取得します。
func (b *Box[T]) Get() T {
	return b.Value
}

func main() {
	intBox := Box[int]{}
	intBox.Set(123)
	fmt.Println(intBox.Get()) // 123

	stringBox := Box[string]{}
	stringBox.Set("Hello Generics")
	fmt.Println(stringBox.Get()) // Hello Generics
}
```

この例では、`Box` 型は任意の型の値を保持することができる汎用コンテナとして定義されています。`Set` メソッドと `Get` メソッドを使用して、`Box` 内の値を設定および取得できます。ジェネリックを使用することで、同じロジックを異なる型に適用できる柔軟なコードを書くことが可能になります。

これらのプログラムは、Go言語がジェネリックをサポートしている環境（ `Go 1.18`  以降）で実行することができます。ジェネリックを使用することで、型安全性を保ちながらコードの再利用性を高めることができます。
