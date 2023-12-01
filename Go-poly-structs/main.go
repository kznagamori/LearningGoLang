package main

import "fmt"

// Greeter は挨拶をするメソッドを持つインターフェースです。
type Greeter interface {
	Greet() string
}

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
