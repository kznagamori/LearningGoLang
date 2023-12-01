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
