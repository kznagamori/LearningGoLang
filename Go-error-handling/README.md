# 一般的なエラー処理

Go言語では、エラー処理は通常、エラーを返り値として関数から返し、呼び出し側でそのエラーを確認して対応するというパターンで行われます。以下に、エラーを返す関数と、その関数を呼び出してエラーを処理するサンプルプログラムを示します。

## サンプルプログラム
この例では、文字列を数値に変換するシンプルな関数を作成し、その関数がエラーを返す場合の処理を示します。

**エラーを返す関数**
```go
package main

import (
    "errors"
    "fmt"
    "strconv"
)

// ConvertToInt は文字列を整数に変換します。
// 変換できない場合はエラーを返します。
func ConvertToInt(str string) (int, error) {
    num, err := strconv.Atoi(str)
    if err != nil {
        return 0, errors.New("変換エラー: " + str + " は整数に変換できません")
    }
    return num, nil
}
```
**エラー処理を行う呼び出し側**
```go
func main() {
    // 正常なケース
    result, err := ConvertToInt("123")
    if err != nil {
        fmt.Println("エラー:", err)
    } else {
        fmt.Println("変換結果:", result)
    }

    // エラーが発生するケース
    result, err = ConvertToInt("abc")
    if err != nil {
        fmt.Println("エラー:", err)
    } else {
        fmt.Println("変換結果:", result)
    }
}
```

このプログラムでは、`ConvertToInt` 関数は文字列を整数に変換しようとします。変換が成功すれば整数と `nil` エラーを返し、失敗すれば0とエラーメッセージを返します。main関数では、この `ConvertToInt`関数を呼び出し、返されたエラーを確認して適切に処理します。

Goでは、このように関数がエラーを返し、呼び出し側でそのエラーを検査することが一般的なエラー処理のパターンです。エラーは値として扱われ、if文を使用してエラーの有無をチェックし、エラーがある場合はそれに応じた処理を行います。

