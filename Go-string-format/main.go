package main

import (
	"fmt"
)

func main() {
	// 整数型のフォーマット
	fmt.Printf("10進数: %d, 2進数: %b, 8進数: %o, 16進数: %x\n", 15, 15, 15, 15)
	fmt.Printf("整数: %d, 符号付き: %+d, 0パディング: %08d\n", 15, 15, 15)

	// 浮動小数点型のフォーマット
	fmt.Printf("浮動小数点: %f, 科学技術記法: %e\n", 123.45678, 123.45678)
	fmt.Printf("浮動小数点: %f, 幅9: %9f, 小数点以下2桁: %.2f, 幅9小数点以下2桁: %9.2f\n", 123.456, 123.456, 123.456, 123.456)

	// 文字列型のフォーマット
	fmt.Printf("文字列: %s, クオート付き: %q, 幅6右詰め: %6s, 幅6左詰め: %-6s\n", "go", "go", "go", "go")
	fmt.Printf("最初の5文字: %.5s, 幅6で最初の2文字: %6.2s\n", "Hello, world!", "Hello, world!")

	// ブール型のフォーマット
	fmt.Printf("ブール: %t\n", true)

	// ポインタのフォーマット
	num := 5
	fmt.Printf("ポインタ: %p\n", &num)
}