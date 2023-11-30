package mypackage

// ファイル内グローバル変数
var fileScopedVariable = "私はfile1.go内でのみアクセス可能です"

// ファイル内グローバル関数
func fileScopedFunction() string {
	return "この関数はfile1.go内でのみ呼び出せます"
}

// ファイル外グローバル変数
var PackageScopedVariable = "私はどのファイルからでもアクセス可能です"

// ファイル外グローバル関数
func PackageScopedFunction() string {
	return "この関数はどのファイルからでも呼び出せます"
}
