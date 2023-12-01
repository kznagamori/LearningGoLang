# 構造体のC#みたいなプロパティを定義、使用

Go言語にはC#のようなプロパティの概念はありませんが、ゲッター（ `Getter` ）とセッター（ `Setter` ）というメソッドを定義することで、プロパティのような機能を模倣することができます。以下に、ゲッターとセッターを使用したサンプルプログラムを示します。

## サンプルプログラム
この例では、`Person` 構造体に `name` というプライベートなフィールドを持たせ、`GetName` と `SetName` というゲッターとセッターを定義しています。

**`Person` 構造体とゲッター・セッター**
```go
package main

import (
    "fmt"
    "strings"
)

// Person 構造体
type Person struct {
    name string // プライベートなフィールド
}

// NewPerson は新しいPersonインスタンスを作成するコンストラクタです
func NewPerson(name string) *Person {
    return &Person{name: name}
}

// GetName はnameフィールドの値を取得します（ゲッター）
func (p *Person) GetName() string {
    return p.name
}

// SetName はnameフィールドの値を設定します（セッター）
func (p *Person) SetName(name string) {
    p.name = strings.TrimSpace(name) // 名前の前後の空白を取り除く
}

func main() {
    // Personインスタンスの作成
    person := NewPerson(" Taro ")

    // ゲッターを使用して名前を取得
    fmt.Println("名前:", person.GetName())

    // セッターを使用して名前を変更
    person.SetName("Hanako")
    fmt.Println("新しい名前:", person.GetName())
}
```

このプログラムでは、`Person` 構造体の `name` フィールドがプライベートであり、外部から直接アクセスすることはできません。代わりに、 `GetName` と `SetName` メソッドを使用して間接的に `name` フィールドの値を取得・設定します。

Goでは、このようにメソッドを使用してデータのカプセル化を行うことが一般的です。C#のプロパティのようにフィールドに直接アクセスする機能はありませんが、ゲッターとセッターを定義することで同様の機能を実現することができます。
