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
