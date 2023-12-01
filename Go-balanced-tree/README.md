# 平衡二分木構造

平衡二分木（バランスの取れたバイナリツリー）を実装するためのサンプルプログラムを示します。この例では、AVLツリー（自己バランス二分探索木）を実装します。AVLツリーは、挿入や削除のたびにツリーのバランスを保つために回転操作を行います。

## サンプルプログラム
```go
package main

import (
	"fmt"
)

// TreeNode はAVLツリーのノードを表します。
type TreeNode struct {
	Value  int
	Height int
	Left   *TreeNode
	Right  *TreeNode
}

// NewTreeNode は新しいTreeNodeを作成します。
func NewTreeNode(value int) *TreeNode {
	return &TreeNode{Value: value, Height: 1}
}

// GetHeight はノードの高さを返します。
func (n *TreeNode) GetHeight() int {
	if n == nil {
		return 0
	}
	return n.Height
}

// GetBalanceFactor はノードのバランス係数を計算します。
func (n *TreeNode) GetBalanceFactor() int {
	if n == nil {
		return 0
	}
	return n.Left.GetHeight() - n.Right.GetHeight()
}

// RightRotate は右回転を行います。
func (n *TreeNode) RightRotate() *TreeNode {
	leftNode := n.Left
	rightOfLeftNode := leftNode.Right

	leftNode.Right = n
	n.Left = rightOfLeftNode

	n.Height = max(n.Left.GetHeight(), n.Right.GetHeight()) + 1
	leftNode.Height = max(leftNode.Left.GetHeight(), leftNode.Right.GetHeight()) + 1

	return leftNode
}

// LeftRotate は左回転を行います。
func (n *TreeNode) LeftRotate() *TreeNode {
	rightNode := n.Right
	leftOfRightNode := rightNode.Left

	rightNode.Left = n
	n.Right = leftOfRightNode

	n.Height = max(n.Left.GetHeight(), n.Right.GetHeight()) + 1
	rightNode.Height = max(rightNode.Left.GetHeight(), rightNode.Right.GetHeight()) + 1

	return rightNode
}

// Insert はツリーに新しい要素を追加します。
func (n *TreeNode) Insert(value int) *TreeNode {
	if n == nil {
		return NewTreeNode(value)
	}

	if value < n.Value {
		n.Left = n.Left.Insert(value)
	} else if value > n.Value {
		n.Right = n.Right.Insert(value)
	} else {
		return n
	}

	n.Height = 1 + max(n.Left.GetHeight(), n.Right.GetHeight())

	balanceFactor := n.GetBalanceFactor()

	// Left Left Case
	if balanceFactor > 1 && value < n.Left.Value {
		return n.RightRotate()
	}

	// Right Right Case
	if balanceFactor < -1 && value > n.Right.Value {
		return n.LeftRotate()
	}

	// Left Right Case
	if balanceFactor > 1 && value > n.Left.Value {
		n.Left = n.Left.LeftRotate()
		return n.RightRotate()
	}

	// Right Left Case
	if balanceFactor < -1 && value < n.Right.Value {
		n.Right = n.Right.RightRotate()
		return n.LeftRotate()
	}

	return n
}

// max は2つの整数の最大値を返します。
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
	root := NewTreeNode(80)
	fmt.Printf("Hight: %d\n", root.Height)
	root = root.Insert(70)
	fmt.Printf("Hight: %d\n", root.Height)
	root = root.Insert(60)
	fmt.Printf("Hight: %d\n", root.Height)
	root = root.Insert(50)
	fmt.Printf("Hight: %d\n", root.Height)
	root = root.Insert(40)
	fmt.Printf("Hight: %d\n", root.Height)
	root = root.Insert(30)
	fmt.Printf("Hight: %d\n", root.Height)
	root = root.Insert(20)
	fmt.Printf("Hight: %d\n", root.Height)

	root.InOrderTraversal()
}
```
