package main

import "fmt"

// Person 構造体の定義
type Person struct {
	Name string
	Age  int
}

// Greet メソッドはPersonのインスタンスに対して挨拶を行います。
func (p Person) Greet() string {
	return fmt.Sprintf("こんにちは、私の名前は%sです。年齢は%d歳です。", p.Name, p.Age)
}

func main() {
	// Person 構造体のインスタンスを作成
	person := Person{Name: "Taro", Age: 30}

	// Greet メソッドを呼び出し
	fmt.Println(person.Greet())
}
