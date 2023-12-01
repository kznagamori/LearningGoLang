package main

import (
	"fmt"
)

func main() {
	// ハッシュテーブル（map）の初期化
	hashTable := make(map[string]int)

	// 要素の挿入
	hashTable["apple"] = 5
	hashTable["banana"] = 10
	hashTable["orange"] = 15

	// 要素のアクセスと表示
	fmt.Println("Apple count:", hashTable["apple"])
	fmt.Println("Banana count:", hashTable["banana"])

	// 要素の存在チェック
	value, exists := hashTable["grape"]
	if exists {
		fmt.Println("Grape count:", value)
	} else {
		fmt.Println("Grape not found")
	}

	// 要素の削除
	delete(hashTable, "banana")

	// ハッシュテーブルの内容を表示
	fmt.Println("Current hash table:", hashTable)
}
