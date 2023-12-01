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
