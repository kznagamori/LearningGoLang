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
