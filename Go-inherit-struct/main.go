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

// Derived は Base 構造体を埋め込んだ派生構造体です
type Derived struct {
	Base   // 匿名フィールドとして Base 構造体を埋め込む
	Degree string
}

// Introduce は Derived 構造体のメソッドです
func (d Derived) Introduce() string {
	return d.Greet() + " 私は" + d.Degree + "の学生です。"
}

func main() {
	// Derived 構造体のインスタンスを作成
	d := Derived{
		Base:   Base{Name: "Taro"},
		Degree: "コンピュータサイエンス",
	}

	// Base の Greet と Derived の Introduce メソッドを呼び出す
	fmt.Println(d.Greet())     // Base のメソッド
	fmt.Println(d.Introduce()) // Derived のメソッド
}
