package main

import "fmt"

func main() {
	fmt.Println("main function started")
	defer fmt.Println("main function deferred call")

	fmt.Println("main function running")

	exampleFunction()
	fmt.Println("main function ending")
}

func exampleFunction() {
	fmt.Println("exampleFunction started")
	defer fmt.Println("exampleFunction deferred call")
	fmt.Println("exampleFunction running")
}
