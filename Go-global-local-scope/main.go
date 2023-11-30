package main

import (
	"Go-global-local-scope/mypackage"
	"fmt"
)

func main() {
	// file1.goのfileScopedVariableとfileScopedFunctionは参照できない
	//fmt.Println(mypackage.fileScopedVariable)
	//fmt.Println(mypackage.fileScopedFunction())

	// file1.goのPackageScopedVariableとPackageScopedFunctionを使用
	fmt.Println(mypackage.PackageScopedVariable)
	fmt.Println(mypackage.PackageScopedFunction())
}
