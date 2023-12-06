package main

import (
	"container/list"
	"fmt"
)

func main() {
	// リンクドリストの作成
	l := list.New()

	// 要素の追加
	l.PushBack(1)  // 末尾に追加
	l.PushFront(2) // 先頭に追加
	l.PushBack(3)  // 末尾に追加

	// リンクドリストの表示
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// 特定の要素の削除
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == 2 {
			l.Remove(e)
			break
		}
	}

	fmt.Println("After removing an element:")

	// 要素削除後のリンクドリストの表示
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
