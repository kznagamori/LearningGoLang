// main.go
package main

import (
	"Go-multi-pack-file/greetings" // ローカルの `greetings` パッケージをインポート
	"fmt"
)

func main() {
	// greetings パッケージの Greet 関数を呼び出し
	greeting := greetings.Greet("Taro")
	fmt.Println(greeting)
}
