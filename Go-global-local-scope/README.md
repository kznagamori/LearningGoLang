# ファイル内グローバル変数と関数、ファイル外グローバル変数と関数

グローバル変数と関数は、そのスコープ（可視性範囲）に基づいて、ファイル内グローバル（同一ファイル内でのみアクセス可能）またはファイル外グローバル（他のファイルからもアクセス可能）として定義されます。

## ファイル内グローバル変数と関数
ファイル内グローバル変数や関数は、その定義を含むファイル内でのみアクセス可能です。これらは通常、小文字で始まる名前を使って定義されます。

### file1.go
```go
package mypackage

// ファイル内グローバル変数
var fileScopedVariable = "私はfile1.go内でのみアクセス可能です"

// ファイル内グローバル関数
func fileScopedFunction() string {
    return "この関数はfile1.go内でのみ呼び出せます"
}
```

## ファイル外グローバル変数と関数
ファイル外グローバル変数や関数は、他のファイルからもアクセス可能です。これらは大文字で始まる名前を使って定義され、Goの公開規則に従います。

### file1.go
```go
package mypackage

// ファイル外グローバル変数
var PackageScopedVariable = "私はどのファイルからでもアクセス可能です"

// ファイル外グローバル関数
func PackageScopedFunction() string {
    return "この関数はどのファイルからでも呼び出せます"
}
```

### main.go
```go
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
```

この例では、`file1.go` に定義された `PackageScopedVariable`と`PackageScopedFunction` は、別のファイル（main.go）からアクセスできます。  
これに対して、`fileScopedVariable` と `fileScopedFunction` は`file1.go` 内でのみアクセス可能です。

Go言語においては、名前が大文字で始まる変数や関数は公開（エクスポートされる）され、他のパッケージからアクセス可能です。一方、小文字で始まる名前は非公開であり、定義されたパッケージ内でのみアクセス可能です。