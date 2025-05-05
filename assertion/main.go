package main

import "fmt"

func main() {
	var v interface{}

	v = 100
	n, ok := v.(int)
	fmt.Println(n, ok)

	s, ok := v.(string)
	fmt.Println(s, ok)
}
