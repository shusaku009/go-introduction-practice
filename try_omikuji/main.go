package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

// おみくじの結果をランダムで出力するための関数を作成
func randomOmikuji() string {
	arr := [...]string{"大吉", "中吉", "小吉", "吉", "凶", "大凶"}
	return "あなたの運勢は" + arr[rand.Intn(len(arr))]
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, randomOmikuji())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
