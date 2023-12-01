# パブリックな構造体のメンバ、メソッドを定義とプライベートな構造体のメンバ、メソッドを定義する

Go言語では、パブリック（公開）な構造体、メンバ、メソッドとプライベート（非公開）な構造体、メンバ、メソッドを定義する際、名前の最初の文字を大文字（パブリック用）または小文字（プライベート用）で始めることによって可視性を制御します。

以下のサンプルプログラムでは、パブリックな `Person` 構造体とそのパブリックメソッド `Greet` 、そしてプライベートな `secret` 構造体とそのプライベートメソッド `whisper` を定義しています。

## サンプルプログラム

### パブリックな構造体とメソッド
**file1.go**

```go
package mypackage

import "fmt"

// Person 構造体はパブリックです（大文字で始まっているため）
type Person struct {
    Name string // パブリックなフィールド
    Age  int    // パブリックなフィールド
}

// Greet メソッドはパブリックです
func (p Person) Greet() string {
    return fmt.Sprintf("こんにちは、私の名前は%sです。年齢は%d歳です。", p.Name, p.Age)
}

```
### プライベートな構造体とメソッド
**file1.go**
```go
// secret 構造体はプライベートです（小文字で始まっているため）
type secret struct {
	thoughts string // プライベートなフィールド
}

// whisper メソッドはプライベートです
func (s secret) whisper() string {
	return fmt.Sprintf("秘密の思い: %s", s.thoughts)
}

// SecretFunc メソッドはパブリックです
func SecretFunc() string {
	// 同じパッケージ内からは呼び出せる
	s := secret{thoughts: "これは秘密です"}
	return s.whisper()
}

```

### 呼び出し
**main.go**
```go
package main

import (
	"Go-public-private/mypackage"
	"fmt"
)

func main() {
	// パブリックなPerson構造体のインスタンスを作成
	p := mypackage.Person{Name: "Taro", Age: 30}
	fmt.Println(p.Greet())

	// プライベートなのでアクセスできない
	// プライベートなsecret構造体のインスタンスを作成
	//s := mypackage.secret{thoughts: "これは秘密です"}
	//fmt.Println(s.whisper())

	// プライベートなsecretの呼び出し
	fmt.Println(mypackage.SecretFunc())

}
```

このプログラムでは、`Person` 構造体とその `Greet` メソッドはどちらもパブリックです。これは、他のパッケージからもアクセス可能であることを意味します。一方で、`secret` 構造体とその `whisper` メソッドはプライベートで、定義されたパッケージ内からのみアクセス可能です。

Goにおいては、構造体やメソッドの可視性は、命名規則によって制御されます。大文字で始まる名前は公開（エクスポート）され、小文字で始まる名前は非公開となります。
