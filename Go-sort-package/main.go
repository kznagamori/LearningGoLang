package main

import (
	"fmt"
	"sort"
)

func main() {
	// 整数のスライスをソート
	intSlice := []int{4, 2, 7, 3, 8}
	sort.Ints(intSlice)
	fmt.Println("Sorted integers:", intSlice)

	// 文字列のスライスをソート
	stringSlice := []string{"banana", "apple", "cherry"}
	sort.Strings(stringSlice)
	fmt.Println("Sorted strings:", stringSlice)

	// カスタム比較関数を使ってスライスをソート
	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 30},
		{"Bob", 25},
		{"Clara", 27},
	}

	// 年齢でソート
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println("Sorted by age:", people)
}
