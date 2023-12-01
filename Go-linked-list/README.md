# リスト構造

Go言語でリスト構造を実装するための簡単なサンプルプログラムを紹介します。この例では、単純な連結リストを定義し、いくつかの要素をリストに追加し、リストを繰り返し処理して要素を表示します。

## サンプルプログラム

```go
package main

import (
	"fmt"
)

// ListNode はリストのノードを表します。
type ListNode struct {
	Value int
	Next  *ListNode
}

// Insert は新しいノードをリストの最後に追加します。
func (n *ListNode) Insert(value int) {
	current := n
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &ListNode{Value: value}
}

// Display はリストの要素を表示します。
func (n *ListNode) Display() {
	current := n
	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.Next
	}
	fmt.Println()
}

func main() {
	// 最初のノードを作成します。
	head := &ListNode{Value: 1}

	// リストに要素を追加します。
	head.Insert(2)
	head.Insert(3)
	head.Insert(4)

	// リストの要素を表示します。
	head.Display()
}

```
このプログラムでは、`ListNode` 構造体を使用してリストの各ノードを表します。`Insert` メソッドは新しいノードをリストの最後に追加し、 `Display` メソッドはリストのすべての要素を表示します。`main` 関数では、リストを作成し、いくつかの要素を追加してから、それらを表示します。

