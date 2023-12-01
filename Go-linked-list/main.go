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
