package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("別のゴールーチン")
	}()

	fmt.Println("メインゴールーチン")
	time.Sleep(50*time.Millisecond)
}