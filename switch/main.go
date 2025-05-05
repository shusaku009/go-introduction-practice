package main

import (
	"fmt"
)

func main() {
	var i interface{}

	i = 100
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "hoge")
	default:
		fmt.Println("default")
	}
}
