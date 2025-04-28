package main

import "fmt"

func f() int {
	n := 100
	return n
}

func g() int {
	n := 200
	return n
}

func main() {
	fmt.Println(f())
	fmt.Println(g())
}
