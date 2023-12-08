package main

import (
	"fmt"
	"strings"
)

// Filter はスライスから条件に合致する要素だけを抽出します。
func Filter[T any](s []T, fn func(T) bool) []T {
	var result []T
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Map はスライスの各要素に関数を適用します。
func Map[T any, U any](s []T, fn func(T) U) []U {
	var result []U
	for _, v := range s {
		result = append(result, fn(v))
	}
	return result
}

func main() {
	// スライスの初期化
	words := []string{"go", "rust", "java", "python"}

	// Filter: 文字列の長さが3以上の要素を抽出
	longWords := Filter(words, func(s string) bool {
		return len(s) > 3
	})
	fmt.Println("長さが3以上の単語:", longWords)

	// Map: 文字列を大文字に変換
	upperWords := Map(words, strings.ToUpper)
	fmt.Println("大文字に変換した単語:", upperWords)
}
