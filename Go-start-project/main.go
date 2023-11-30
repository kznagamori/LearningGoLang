package main

import (
	"fmt"
	"os"
)

func main() {
	// コマンドライン引数を取得
	args := os.Args[1:]

	// 引数がない場合は、エラーメッセージを表示
	if len(args) == 0 {
		fmt.Println("名前を入力してください")
		os.Exit(1)
	}

	// 挨拶を表示
	name := args[0]
	fmt.Printf("こんにちは、%sさん!\n", name)
}
