package main

import (
	"fmt"
	"go-introduction-practice/package/greeting"
)

var msg string = "hello"

func f() { println(msg) }

func main() {
	txt := greeting.Do()
	fmt.Println(txt)

	f()
	msg = "hi, gophers"
	f()
}
