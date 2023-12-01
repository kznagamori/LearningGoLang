package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// リングバッファのサイズを指定して初期化
	r := ring.New(5)

	// リングバッファにデータを追加
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}

	// リングバッファの要素を表示
	r.Do(func(p interface{}) {
		fmt.Println(p)
	})

	// 特定の要素を操作
	r.Move(2) // 2つ先に移動
	r.Value = 100

	fmt.Println("After modification:")

	// 変更後のリングバッファの要素を表示
	r.Do(func(p interface{}) {
		fmt.Println(p)
	})
}
