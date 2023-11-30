// greetings/greetings.go
package greetings

import "fmt"

// Greet は名前を受け取り、挨拶文を返します。
func Greet(name string) string {
	return fmt.Sprintf("こんにちは、%sさん!", name)
}
