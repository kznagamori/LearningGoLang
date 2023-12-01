# 二分木構造

Go言語で二分木（バイナリツリー）構造を実装するためのサンプルプログラムを以下に示します。この例では、簡単なバイナリツリーを定義し、要素をツリーに追加（挿入）し、ツリーを順次探索（インオーダー・トラバーサル）して要素を表示する基本的な操作を実行します。

## サンプルプログラム
```go
package main

import (
	"fmt"
)

// TreeNode はバイナリツリーのノードを表します。
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// Insert はツリーに新しい要素を追加します。
func (n *TreeNode) Insert(value int) {
	if value <= n.Value {
		if n.Left == nil {
			n.Left = &TreeNode{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

// InOrderTraversal はツリーをインオーダーで探索します。
func (n *TreeNode) InOrderTraversal() {
	if n != nil {
		n.Left.InOrderTraversal()
		fmt.Printf("%d ", n.Value)
		n.Right.InOrderTraversal()
	}
}

func main() {
	// ツリーの根を作成します。
	root := &TreeNode{Value: 5}

	// ツリーに要素を追加します。
	root.Insert(3)
	root.Insert(7)
	root.Insert(2)
	root.Insert(4)
	root.Insert(6)
	root.Insert(8)

	// ツリーをインオーダーで探索し、要素を表示します。
	root.InOrderTraversal()
}

```
