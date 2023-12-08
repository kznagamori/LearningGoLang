package main

import "fmt"

// Map はスライスの各要素に関数を適用します。
func Map[T any, U any](s []T, fn func(T) U) []U {
	var result []U
	for _, v := range s {
		result = append(result, fn(v))
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	// Map 関数にラムダ式（無名関数）を渡して、各要素を2倍にする
	doubled := Map(nums, func(n int) int {
		return n * 2
	})

	fmt.Println(doubled) // [2 4 6 8 10]

	// 即時関数（IIFE）を定義し実行
	sum := func(nums ...int) int {
		total := 0
		for _, n := range nums {
			total += n
		}
		return total
	}(1, 2, 3, 4, 5)

	fmt.Println(sum) // 15
}
