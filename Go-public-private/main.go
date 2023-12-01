package main

import (
	"Go-public-private/mypackage"
	"fmt"
)

func main() {
	// パブリックなPerson構造体のインスタンスを作成
	p := mypackage.Person{Name: "Taro", Age: 30}
	fmt.Println(p.Greet())

	// プライベートなのでアクセスできない
	// プライベートなsecret構造体のインスタンスを作成
	//s := mypackage.secret{thoughts: "これは秘密です"}
	//fmt.Println(s.whisper())

	// プライベートなsecretの呼び出し
	fmt.Println(mypackage.SecretFunc())

}
