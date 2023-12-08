package main

import "fmt"

// Compare は2つの値を比較し、等しいかどうかを返します。
// この関数は任意の型（T）に対して動作します。
func Compare[T comparable](a, b T) bool {
	return a == b
}

func main() {
	fmt.Println(Compare[int](5, 5))            // true
	fmt.Println(Compare[string]("go", "rust")) // false
}
