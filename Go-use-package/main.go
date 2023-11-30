package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// ルーターの初期化
	r := mux.NewRouter()

	// エンドポイントとハンドラーの設定
	r.HandleFunc("/", HomeHandler)

	// サーバーの起動
	log.Fatal(http.ListenAndServe(":8000", r))
}

// ハンドラー関数
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
