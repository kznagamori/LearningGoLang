// greetings/greetings.go
package greetings

import "fmt"

// Greet は挨拶を返します。
func Greet(name string) string {
	return fmt.Sprintf("こんにちは、%sさん!", name)
}
