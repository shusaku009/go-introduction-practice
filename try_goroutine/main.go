package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("メイン関数を終了します")

	go func() {
		defer fmt.Println("goroutine1 done")
		time.Sleep(1 * time.Second)
	}()
	go func() {
		defer fmt.Println("goroutine2 done")
		time.Sleep(4 * time.Second)
	}()
	time.Sleep(5 * time.Second)
}